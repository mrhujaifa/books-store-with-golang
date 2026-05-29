package auth

import (
	"context"
	"fmt"
	"strings"
)

// CustomClaims contains custom data we want to parse from the JWT.
type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	if c.Scope == "" {
		return nil // No scope is valid - not all endpoints require permissions
	}

	if strings.TrimSpace(c.Scope) != c.Scope {
		return fmt.Errorf("scope claim has invalid whitespace")
	}

	if strings.Contains(c.Scope, "  ") {
		return fmt.Errorf("scope claim contains double spaces")
	}

	return nil
}
