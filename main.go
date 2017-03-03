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
	"strconv"
)

var vidDir = "G:/Videos"
var thumbCache = "" //TODO: this
var staticDir = "C:/Users/mdumf_000/Brogramming/src/github.com/Squaar/Webm-Gallery/static"
var port = 8080

type logHandler func(http.ResponseWriter, *http.Request)

//TODO: command line args
func main() {
	http.Handle("/", logHandler(static))
    http.Handle("/files", logHandler(files))
    http.Handle("/file", logHandler(file))
    http.Handle("/file/", logHandler(file))

    log.Println("Listening on port " + strconv.Itoa(port))
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func (fn logHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	log.Println(r.RemoteAddr + " - " + r.URL.String())
	fn(rw, r)
}

func static(rw http.ResponseWriter, r *http.Request){
	if r.URL.String() == "/" || r.URL.String() == "/gallery"{
		log.Println("Redirecting to /gallery.html")
		http.Redirect(rw, r, "/gallery.html", 303)
		return
	}

	filePath := filepath.Join(staticDir, r.URL.String())

	err := sendFile(rw, filePath)
	if err != nil{
		log.Println(err)
	}
}

//return newest first?
//TODO: don't fatal log
func files(rw http.ResponseWriter, r *http.Request){
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
	var fileName string
	if strings.ContainsRune(r.URL.String(), '?'){
		fileName = r.URL.Query().Get("file")
	} else{
		_, fileName = filepath.Split(r.URL.String())
	}
	fullPath := filepath.Join(vidDir, fileName)

	err := sendFile(rw, fullPath)
	if err != nil{
		log.Println(err)
	}
}

//TODO: read in chunks for safety?
//TODO: look up how to panic and maybe do it here
func sendFile(rw http.ResponseWriter, filePath string) error{
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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