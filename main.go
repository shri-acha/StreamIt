package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"StreamIt/handler"
)
//my shit got deleted so had to generate this,might f'up somewhere just keep on checking
func main() {
	// Serve video files
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./uploads"))))

	// Serve static files (including index.html)
	fs := http.FileServer(http.Dir("./static"))
	
	http.Handle("/static/",http.StripPrefix("/static/",fs))

	http.HandleFunc("/upload",handler.UploadPageHandler)	
	
	http.HandleFunc("/",handler.HomeHandler)

	// Handle video uploads
	http.HandleFunc("/uploader", handler.UploadHandler)

	// Handle video streaming
	http.HandleFunc("/watch",handler.WatchHandler)

	// Determine the directory of the executable
	exePath, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server...\n")
	fmt.Printf("Serving videos from: %s/videos\n", exePath)
	fmt.Printf("Serving static files from: %s/static\n", exePath)
	fmt.Printf("Server is running on http://localhost:8080\n")

	// Start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
