package handler

import (
	"log"
	"net/http"
	"text/template"
)

func UploadPageHandler(w http.ResponseWriter, r* http.Request){
			
	tmpl,err := template.ParseFiles("./static/index.html")
	if err!=nil{
		http.Error(w,"Server Error",101)
		log.Println(err)
	}

	tmpl.Execute(w,nil)
}
