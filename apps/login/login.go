package login

import (
	"net/http"
	"fmt"
	"github.com/ralphte/chance/apps/cookie"
)


func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	name := r.FormValue("username")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		cookie.SetSession(name, w)
		redirectTarget = "/profile"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

