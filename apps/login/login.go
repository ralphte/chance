package login

import (
        "net/http"
        "fmt"
        "github.com/ralphte/chance/apps/cookie"
        "github.com/ralphte/chance/apps/sql"
        "github.com/ralphte/chance/apps/passhash"
)


func Login(w http.ResponseWriter, r *http.Request) {
        fmt.Println("method:", r.Method) //get request method
        name := r.FormValue("username")
        pass := r.FormValue("password")
        redirectTarget := "/"
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
        }
}
