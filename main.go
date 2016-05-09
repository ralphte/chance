package main

import (
	"fmt"
	"log"
	"net/http"
	"./login"
	"strings"

	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}



func profile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Working Profile!") // write data to response
}

func main() {
	http.HandleFunc("/", login.Login) // setting router rule
	http.HandleFunc("/profile", profile)
	fs := justFilesFilesystem{http.Dir("resources/")}
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(fs)))
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
