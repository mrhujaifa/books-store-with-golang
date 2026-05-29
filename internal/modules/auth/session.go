package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	AppSessionName   = "bookshop_session"
	StateSessionName = "bookshop_oauth_state"
)

func NewSessionStore(secret string) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(secret))

	// gorilla/sessions docs অনুযায়ী cookie options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false, // production এ true
		SameSite: http.SameSiteLaxMode,
	}

	return store
}
