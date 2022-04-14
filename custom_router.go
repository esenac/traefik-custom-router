// Package traefik_custom_router a demo plugin.
package traefik_custom_router

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Public  string `json:"public,omitempty"`
	Private string `json:"private,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Public:  "",
		Private: "",
	}
}

// UserRedirect a UserRedirect plugin.
type UserRedirect struct {
	next   http.Handler
	name   string
	config *Config
}

// New created a new UserRedirect plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &UserRedirect{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (u *UserRedirect) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	authorization := req.Header.Get("Authorization")

	if authorization != "" {
		http.Redirect(rw, req, u.config.Private, http.StatusSeeOther)
		return
	}
	http.Redirect(rw, req, u.config.Public, http.StatusSeeOther)
}
