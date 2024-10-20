package handler

import (
	// "log"
	"net/http"
	"strings"

	// "strings"
	"text/template"
)

func WatchHandler(w http.ResponseWriter,r* http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Invalid Method",http.StatusMethodNotAllowed)
		return
	}

	videoName := r.URL.Query().Get("v")
	if len(videoName) <= 0{
		return
	}
	vnameDirectory := strings.Split(videoName, ".") [0]
	// vnameExtension := strings.Split(videoName, ".") [1]
	// finalFileDirectory := "./uploads/"+vnameDirectory+"/"

	tmpl,err := template.ParseFiles("./static/watch.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := struct {
		VideoCode string;
	}{
		vnameDirectory,
	}

	err = tmpl.Execute(w,data)

	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	return 
}
