package logout
import (
	"net/http"
	"github.com/ralphte/chance/apps/cookie"
)

func Logout(w http.ResponseWriter, request *http.Request) {
	cookie.ClearSession(w)
	http.Redirect(w, request, "/", 302)
}
