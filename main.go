package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"path/filepath"
	"io"
	"strings"
	// "github.com/gorilla/mux"
	"strconv"
)

var vidDir = "G:/Videos"
var staticDir = "C:/Users/mdumf_000/Brogramming/src/github.com/Squaar/Webm-Gallery/static"
var port = 8080

//TODO: command line args
//TODO: fix router
func main() {
	// router := mux.NewRouter()
 //    router.HandleFunc("/files", files)
 //    router.HandleFunc("/file/{file}", file)
 //    router.HandleFunc("/file", file)
 //    http.Handle("/", router)

	http.HandleFunc("/", static)
    http.HandleFunc("/files", files)
    http.HandleFunc("/file", file)	
    http.HandleFunc("/file/", file)

    log.Println("Listening on port " + strconv.Itoa(port))
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func static(rw http.ResponseWriter, r *http.Request){
	log.Println("New Request: " + r.URL.String())
	filePath := filepath.Join(staticDir, r.URL.String())

	err := sendFile (rw, filePath)
	if err != nil{
		log.Println(err)
	}
}

//return newest first?
func files(rw http.ResponseWriter, r *http.Request){
	log.Println("New Request: " + r.URL.String())
	files, err := ioutil.ReadDir(vidDir)
    if err != nil {
    	log.Fatal(err)
    }

    var fileNames []string
    for _, file := range files{
    	fileNames = append(fileNames, file.Name())
    }

    j, err := json.Marshal(fileNames)
    if err != nil{
    	log.Fatal(err)
    }
    rw.Write(j)
}

func file(rw http.ResponseWriter, r *http.Request){
	log.Println("New Request: " + r.URL.String())
	var fileName string
	if strings.ContainsRune(r.URL.String(), '?'){
		fileName = r.URL.Query().Get("file")
	} else{
		_, fileName = filepath.Split(r.URL.String())
	}
	fullPath := filepath.Join(vidDir, fileName)

	err := sendFile (rw, fullPath)
	if err != nil{
		log.Println(err)
	}
}

//TODO: read in chunks for safety?
func sendFile(rw http.ResponseWriter, filePath string) error{
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return err
	}

	stats, err := file.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, stats.Size())
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF{
			return err
	}

	rw.Write(buffer)
	return nil
}