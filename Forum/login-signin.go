package Forum

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("inserer clé")
	store = sessions.NewCookieStore(key)
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func HandleLogout(w http.ResponseWriter, r *http.Request) {

}

func HandleHome(w http.ResponseWriter, r *http.Request) {

}
