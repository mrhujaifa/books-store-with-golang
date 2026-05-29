package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

type Handler struct {
	service Service
	oidc    *OIDCManager
	store   *sessions.CookieStore
	appURL  string
}

func NewHandler(service Service, oidc *OIDCManager, store *sessions.CookieStore, appURL string) *Handler {
	return &Handler{
		service: service,
		oidc:    oidc,
		store:   store,
		appURL:  appURL,
	}
}

func generateState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Step 1: user login শুরু করবে
// Auth0 login page-এ redirect যাবে
func (h *Handler) Login(res http.ResponseWriter, req *http.Request) {
	state, err := generateState()
	if err != nil {
		http.Error(res, "failed to generate state", http.StatusInternalServerError)
		return
	}

	// OAuth state session এ save
	session, _ := h.store.Get(req, StateSessionName)
	session.Values["state"] = state
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	_ = session.Save(req, res)

	// OIDC/OAuth authorization URL তৈরি
	url := h.oidc.OAuth2Config.AuthCodeURL(state)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// Step 1b: signup page-এ redirect
func (h *Handler) Signup(res http.ResponseWriter, req *http.Request) {
	state, err := generateState()
	if err != nil {
		http.Error(res, "failed to generate state", http.StatusInternalServerError)
		return
	}

	session, _ := h.store.Get(req, StateSessionName)
	session.Values["state"] = state
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   300,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	_ = session.Save(req, res)

	url := h.oidc.OAuth2Config.AuthCodeURL(state) + "&screen_hint=signup"
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// Step 2: Auth0 callback
// code -> token exchange -> ID token verify -> claims parse -> DB sync -> cookie set
func (h *Handler) Callback(res http.ResponseWriter, req *http.Request) {
	stateSession, _ := h.store.Get(req, StateSessionName)
	expectedState, _ := stateSession.Values["state"].(string)
	gotState := req.URL.Query().Get("state")

	if expectedState == "" || gotState == "" || expectedState != gotState {
		http.Error(res, "invalid oauth state", http.StatusBadRequest)
		return
	}

	code := req.URL.Query().Get("code")
	if code == "" {
		http.Error(res, "missing authorization code", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// code exchange
	token, err := h.oidc.OAuth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(res, "failed to exchange code", http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(res, "missing id token", http.StatusInternalServerError)
		return
	}

	// ID token verify
	idToken, err := h.oidc.Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(res, "failed to verify id token", http.StatusUnauthorized)
		return
	}

	// claims parse
	var claims Auth0UserClaims
	if err := idToken.Claims(&claims); err != nil {
		http.Error(res, "failed to parse claims", http.StatusInternalServerError)
		return
	}

	// DB sync
	user, err := h.service.SyncUserFromAuth0(claims)
	if err != nil {
		http.Error(res, "failed to sync user", http.StatusInternalServerError)
		return
	}

	// নিজের app session cookie set
	appSession, _ := h.store.Get(req, AppSessionName)
	appSession.Values["user_id"] = strconv.FormatUint(uint64(user.ID), 10)
	appSession.Values["auth0_sub"] = user.Auth0ID
	appSession.Values["email"] = user.Email
	appSession.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	_ = appSession.Save(req, res)

	// state session clear
	stateSession.Options.MaxAge = -1
	_ = stateSession.Save(req, res)

	http.Redirect(res, req, h.appURL, http.StatusTemporaryRedirect)
}

func (h *Handler) Me(res http.ResponseWriter, req *http.Request) {
	session, _ := h.store.Get(req, AppSessionName)
	rawUserID, ok := session.Values["user_id"].(string)
	if !ok || rawUserID == "" {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(map[string]any{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	idUint64, err := strconv.ParseUint(rawUserID, 10, 64)
	if err != nil {
		http.Error(res, "invalid session", http.StatusUnauthorized)
		return
	}

	user, err := h.service.GetUserByID(uint(idUint64))
	if err != nil {
		http.Error(res, "user not found", http.StatusUnauthorized)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]any{
		"success": true,
		"data":    user,
	})
}

func (h *Handler) Logout(res http.ResponseWriter, req *http.Request) {
	session, _ := h.store.Get(req, AppSessionName)
	session.Options.MaxAge = -1
	_ = session.Save(req, res)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]any{
		"success": true,
		"message": "Logged out successfully",
	})
}
