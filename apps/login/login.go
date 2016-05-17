package login

import (
        "net/http"
        "fmt"
        "github.com/ralphte/chance/apps/cookie"
        "github.com/ralphte/chance/apps/sql"
        "github.com/ralphte/chance/apps/passhash"
        "github.com/ralphte/chance/apps/csrf"
)


func Login(w http.ResponseWriter, r *http.Request) {
        fmt.Println("method:", r.Method) //get request method
        name := r.FormValue("username")
        pass := r.FormValue("password")
        token := r.FormValue("token")
        redirectTarget := "/"
        csrfToken := csrf.GetToken(r)
        if csrfToken == "" {
                fmt.Println ("Bad CSRF Token")
                http.Redirect(w, r, redirectTarget, 302)
        }else if csrfToken != token {
                fmt.Println ("Bad CSRF Token")
                http.Redirect(w, r, redirectTarget, 302)
        }else{
                fmt.Println ("Good Token")
                if name != "" && pass != "" {
                        hash := sql.GetUserPassword(name)
                        if passhash.MatchString(hash, pass) == true {
                                fmt.Println("Good Password")
                                cookie.SetSession(name, w)
                                redirectTarget = "/profile"
                                http.Redirect(w, r, redirectTarget, 302)
                        }else{
                                fmt.Println ("Bad Password")
                                http.Redirect(w, r, redirectTarget, 302)
                        }

                }else{
                        http.Redirect(w, r, redirectTarget, 302)
                }
        }
}
