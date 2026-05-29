package config

import "fmt"

type AuthConfig struct {
	Domain   string
	Audience string
}

func LoadAuthConfig(env *Config) (*AuthConfig, error) {
	domain := env.Auth0Domain
	audience := env.Auth0AUDIENCE

	fmt.Println("auth domain from env:", domain)

	return &AuthConfig{
		Domain:   domain,
		Audience: audience,
	}, nil
}
