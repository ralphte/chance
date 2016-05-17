package csrf

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"crypto/rand"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetToken(request *http.Request) (token string) {
	if cookie, err := request.Cookie("csrf_token"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("csrf_token", cookie.Value, &cookieValue); err == nil {
			token = cookieValue["csrf_token"]
		}
	}
	return token
}

func SetToken(token string, response http.ResponseWriter) {
	value := map[string]string{
		"csrf_token": token,
	}
	if encoded, err := cookieHandler.Encode("csrf_token", value); err == nil {
		cookie := &http.Cookie{
			Name:  "csrf_token",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearToken(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "csrf_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// Generate a token
// Source: https://devpy.wordpress.com/2013/10/24/create-random-string-in-golang/
func GenerateToken(length int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return string(bytes)
}