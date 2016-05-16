package profile

import (
	"net/http"
	"html/template"
	"github.com/ralphte/chance/apps/cookie"
	"github.com/ralphte/chance/apps/sql"
	"fmt"
)

type Page struct {
	Username string
	Firstname   string
	Lastname string
	Email   string
}



func Profile(w http.ResponseWriter, r *http.Request) {
	userName := cookie.GetUserName(r)
	if userName != "" {
		t, _ := template.ParseFiles("template/profile.html")
		email, firstname, lastname  := sql.GetUserData(userName)
		profile := Page{
			Username: userName,
			Firstname:   firstname,
			Lastname: lastname,
			Email:   email,
		}
		t.Execute(w, profile)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	userName := cookie.GetUserName(r)
	if userName != "" {
		if firstname != "" && lastname != "" && email != "" {
			sql.UpdateUserData(userName, firstname, lastname, email)
			http.Redirect(w, r, "/profile", 302)
		} else {
			http.Redirect(w, r, "/profile", 302)
		}

	} else {
		http.Redirect(w, r, "/", 302)
	}
}