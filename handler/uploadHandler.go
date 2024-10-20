package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter,r* http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Method Not Allowed!",http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10>>20)
	if err != nil{
		http.Error(w,"Invalid File Size!",http.StatusInsufficientStorage)
		return
}
	f,handler,err := r.FormFile("video")
	
	if err!= nil {
		http.Error(w,"Invalid File Size!",http.StatusInsufficientStorage)
		return
	}

	defer f.Close()

	osFile,err := os.Create("./uploads/"+handler.Filename)	
		
	if err!= nil {
		http.Error(w,"Error Uploading File!",404)
		return
	}
	defer osFile.Close()
	_,err = io.Copy(osFile,f)

	if err != nil{
		http.Error(w,"Error Copying File!",404)
		return
	}

	fmt.Fprintf(w,"File Uploaded Successfully!\n")
	transcodeToHLS(osFile.Name())
	fmt.Fprintf(w,"File Transcoded Successfully!")
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
	return 
}

func transcodeToHLS(inputFile string) error {
	inputFileDirectory := strings.Split(inputFile, ".")
  outputFile := filepath.Join(inputFileDirectory[1], filepath.Base(inputFileDirectory[1])+".m3u8")
	err := os.MkdirAll("."+inputFileDirectory[1],0777)
	
	if err != nil {
		log.Printf("Error has occured!: %v",err)
		return nil
	}

	log.Println(inputFileDirectory[1])
	log.Println(outputFile)
		cmd := exec.Command("ffmpeg",
		      "-i", inputFile,
		      "-profile:v", "baseline",
		      "-level", "3.0",
		      "-start_number", "0",
		      "-hls_time", "10",
		      "-hls_list_size", "0",
		      "-f", "hls",
		      "."+outputFile)

		  output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("ffmpeg error: %v, output: %s", err, string(output))
    }
    return nil
}
