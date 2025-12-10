package requests

import (
	"io"
	"log"
	"net/http"
)

type AddRequest struct {
    Login string `json:"login"`
    Pass  string `json:"pass"`
    Group string `json:"group"`
}

// add new user to database (/etc/ocserv/ocpsasswd)
func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error read request body", r.RemoteAddr)
			http.Error(w, `{"status":"error"}`, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		user := r.Header.Get("user")
		group := r.Header.Get("group")
		log.Println("request / from " + r.RemoteAddr + " user " + user + " body " + string(body) + " group " + group)

		// todo parse request
		// encrypt password
		// put line into ocpasswd

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))

		return
	}

	log.Println("Wrong method", r.RemoteAddr)
	http.Error(w, `{"status":"method"}`, http.StatusInternalServerError)
}
