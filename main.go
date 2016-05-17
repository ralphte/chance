package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/ralphte/chance/apps/profile"
	"github.com/ralphte/chance/apps/login"
	"github.com/ralphte/chance/apps/logout"
	"github.com/ralphte/chance/apps/csrf"
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


func indexPage(w http.ResponseWriter, r *http.Request) {
	token := csrf.GenerateToken(32)
	csrfToken := csrf.GetToken(r)
	if csrfToken == "" {
		csrf.SetToken(token, w)
	}else{
		csrf.ClearToken(w)
		csrf.SetToken(token, w)
	}
	t, _ := template.ParseFiles("template/login.html")
	t.Execute(w, token)
}


var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexPage)
	router.HandleFunc("/profile",profile.Profile).Methods("GET")
	router.HandleFunc("/login",login.Login).Methods("POST")
	router.HandleFunc("/logout",logout.Logout).Methods("POST")
	router.HandleFunc("/update",profile.Update).Methods("POST")
	http.Handle("/", router)
	fs := justFilesFilesystem{http.Dir("resources/")}
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(fs)))
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}