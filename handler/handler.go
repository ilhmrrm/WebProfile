package handler

import (
	"WebProfile/entity"
	"html/template"
	"log"
	"net/http"
	"path"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	index := path.Join("views", "index.html")
	layout := path.Join("views", "layout.html")

	tmpl, err := template.ParseFiles(index, layout)

	if err != nil {
		log.Fatal(err)
	}

	dataEntity := entity.Portofolio{
		Name:  "Ilham Ramadhan",
		Telp:  "087775783217",
		Email: "ilham26ramadhan@gmail.com",
		Title: "My Portofolio",
	}

	err = tmpl.Execute(w, dataEntity)
	if err != nil {
		log.Fatal(err)
	}
}
