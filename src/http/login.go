package http

import (
	"net/http"
	"time"

	"github.com/crossphoton/drive-encrypt/src"
	"github.com/google/uuid"
)

var allowed_tokens map[string]bool = make(map[string]bool)
var global_password string = ""

func Login(w http.ResponseWriter, r *http.Request) {
	_, password, ok := r.BasicAuth()
	if ok {
		err := src.CheckPassword(password)
		if err == nil {
			if global_password == "" {
				global_password = password
			}
			token := uuid.NewString()
			http.SetCookie(w, &http.Cookie{
				Name:   AUTH_COOKIE_NAME,
				Value:  token,
				Secure: true,
				MaxAge: int(AUTH_TIMEOUT),
			})
			w.WriteHeader(http.StatusOK)
			allowed_tokens[token] = true

			go func() {
				time.Sleep(AUTH_TIMEOUT)
				delete(allowed_tokens, token)
			}()
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie(AUTH_COOKIE_NAME)
		if err != nil && cookie.Valid() != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		allowed, ok := allowed_tokens[cookie.Value]
		if ok && allowed {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
