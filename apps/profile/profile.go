package profile

import (
	"net/http"
	"html/template"
	"github.com/ralphte/chance/apps/cookie"
)


func Profile(w http.ResponseWriter, r *http.Request) {
	userName := cookie.GetUserName(r)
	if userName != "" {
		t, _ := template.ParseFiles("template/profile.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", 302)
	}



}

