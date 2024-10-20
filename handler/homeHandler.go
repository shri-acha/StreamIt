package handler

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r* http.Request){
	tmpl,err := template.ParseFiles("./static/home.html")
	
	if err != nil {
		http.Error(w,"Error loading file!",http.StatusNotImplemented)
		log.Println(err)
		return 
	}

	videoNames := listFolders()

	vidData := struct {
		VideoData []string
	}{
		videoNames,
	}

	tmpl.Execute(w,vidData);
}

func listFolders() []string{
	var folderArray []string;
	dirList,err := os.ReadDir("./uploads")
	
	if err != nil{
		log.Println(err)
		return nil
	}
	for _,v := range dirList{
		if v.IsDir(){
			folderArray = append(folderArray,v.Name())
		}
	}
	return folderArray
}
