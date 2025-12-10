package requests

import (
	"html/template"
	"log"
	"net/http"
	"ocserv-admin/internal/model"
	"ocserv-admin/internal/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	user := r.Header.Get("user")
	group := r.Header.Get("group")
	
	log.Println("request / from " + r.RemoteAddr + " user " + user)

	base, online, err := utils.UpdateData()
	if err != nil {
		log.Println("Error load data "+err.Error(), r.RemoteAddr)
		http.Error(w, "Error load data", http.StatusInternalServerError)
		return
	}

	data := model.PageData{}

	if group == "" {
		data.BaseRecords = base
		data.OnlineRecords = online
	} else {
		b := []model.Base{}
		for _, r := range base {
			if r.Group == group {
				b = append(b, r)
			}
		}
		data.BaseRecords = b

		o := []model.Online{}
		for _, r := range online {
			if utils.ExistGroup(&b, r.User, group) {
				o = append(o, r)
			}
		}
		data.OnlineRecords = o
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("Error load html template "+err.Error(), r.RemoteAddr)
		http.Error(w, "Error load html template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error parse html template "+err.Error(), r.RemoteAddr)
		http.Error(w, "Error parse html template", http.StatusInternalServerError)
		return
	}
}
