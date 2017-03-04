package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var logFile = "/home/Squaar/brogramming/logs/gallery.log"
var vidDir = "/mnt/g/Videos"
var thumbCache = "" //TODO: this
var staticDir = "/home/Squaar/brogramming/go/src/github.com/Squaar/Webm-Gallery/static"
var port = 8080

type logHandler func(http.ResponseWriter, *http.Request)

//TODO: command line args
func main() {
	openLog(logFile)

	http.Handle("/", logHandler(static))
	http.Handle("/files", logHandler(files))
	http.Handle("/file", logHandler(file))
	http.Handle("/file/", logHandler(file))
	http.Handle("/thumb/", logHandler(thumb))

	log.Println("Listening on port " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func openLog(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func (fn logHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " - " + r.URL.String())
	fn(rw, r)
}

func static(rw http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" || r.URL.String() == "/gallery" {
		log.Println("Redirecting to /gallery.html")
		http.Redirect(rw, r, "/gallery.html", 303)
		return
	}

	filePath := filepath.Join(staticDir, r.URL.String())
	http.ServeFile(rw, r, filePath)
}

//return newest first?
//TODO: don't fatal log
func files(rw http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(vidDir)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	j, err := json.Marshal(fileNames)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write(j)
}

func file(rw http.ResponseWriter, r *http.Request) {
	var fileName string
	if strings.ContainsRune(r.URL.String(), '?') {
		fileName = r.URL.Query().Get("file")
	} else {
		_, fileName = filepath.Split(r.URL.String())
	}

	fullPath := filepath.Join(vidDir, fileName)
	http.ServeFile(rw, r, fullPath)
}

func thumb(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, staticDir+"/img/SpicyDancer.png")
}
