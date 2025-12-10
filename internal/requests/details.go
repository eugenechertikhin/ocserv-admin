package requests

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// send details about logged user
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		user := r.Header.Get("user")

		vars := mux.Vars(r)
		id := vars["id"]

		log.Println("request /details from " + r.RemoteAddr + " user " + user + " id " + id)

		// 	SHOW_DETAILS = "occtl show id %d"

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))

		return
	}

	log.Println("Wrong method", r.RemoteAddr)
	http.Error(w, `{"status":"method"}`, http.StatusInternalServerError)
}
