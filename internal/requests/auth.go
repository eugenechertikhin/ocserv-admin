package requests

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"ocserv-admin/internal/model"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", "Basic realm=ocserv-admin")
			http.Error(w, "Authorization required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(auth, "Basic ") {
			http.Error(w, "Invalid authorization header", http.StatusBadRequest)
			return
		}

		cred := strings.TrimPrefix(auth, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(cred)
		if err != nil {
			http.Error(w, "Invalid authorization header", http.StatusBadRequest)
			return
		}

		pair := strings.SplitN(string(decoded), ":", 2)
		if len(pair) >= 2 {
			for _, b := range model.Basic {
				if pair[0] == b.User && pair[1] == b.Password {
					r.Header.Add("user", b.User)
					r.Header.Add("password", b.Password)
					r.Header.Add("group", b.Group)
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("Unauthorized request ", decoded)
	}
}
