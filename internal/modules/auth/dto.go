package auth

// ID token claims থেকে এই data parse করবো.
type Auth0UserClaims struct {
	Sub     string `json:"sub"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
