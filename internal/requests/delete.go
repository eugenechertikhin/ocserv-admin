package requests

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Delete user from database (/etc/ocserv/ocpasswd)
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		user := r.Header.Get("user")

		vars := mux.Vars(r)
		login := vars["login"]

		log.Println("request /delete from " + r.RemoteAddr + " user " + user + " login " + login)

		// todo delete from ocpasswd

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))

		return
	}

	log.Println("Wrong method", r.RemoteAddr)
	http.Error(w, `{"status":"method"}`, http.StatusInternalServerError)
}
