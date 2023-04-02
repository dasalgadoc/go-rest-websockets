package domain

import "errors"

type Config struct {
	Port      string
	JWTSecret string
	Database  string
}

func (c Config) ConfigErrors() error {
	if c.Port == "" {
		return errors.New("port is required")
	}

	if c.JWTSecret == "" {
		return errors.New("jwt secret is required")
	}

	if c.Database == "" {
		return errors.New("secret is required")
	}

	return nil
}
