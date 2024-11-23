package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/u2takey/ffmpeg-go"
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
	
	fileType := strings.Split(handler.Filename, ".")[1]
	
	if !validateFileType(fileType){
		http.Error(w,"Wrong Filetype!",http.StatusUnsupportedMediaType)
		return
	}


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
	go func() {
		err := transcodeToHLS(osFile.Name())
		if err != nil {
			log.Println("Error in transcoding:", err)
		} else {
			log.Println("File transcoded successfully:")
		}
	}()

	http.Redirect(w,r,"/",http.StatusSeeOther)
	
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
		// cmd := exec.Command("ffmpeg",
		//       "-i", inputFile,
		//       "-profile:v", "baseline",
		//       "-level", "3.0",
		//       "-start_number", "0",
		//       "-hls_time", "10",
		//       "-hls_list_size", "0",
		//       "-f", "hls",
		//       "."+outputFile)
		//
		//   output, err := cmd.CombinedOutput()
		err = ffmpeg_go.Input(inputFile).
						Output("."+outputFile,
						ffmpeg_go.KwArgs{"profile:v":"baseline",
						"level":"3.0","start_number":"0",
						"hls_time":"10","hls_list_size":"0","f":"hls",
					}).Run();	


    if err != nil {
        return fmt.Errorf("ffmpeg error: %v", err)
    }
    return nil
}

func validateFileType(fileType string) bool{

	fileTypes := []string{"mkv","mp4"}
	
	for _,f := range fileTypes{
		if f==fileType {
			return true
		}
	}
	return false
}
