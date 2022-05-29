package Forum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("inserer clé")
	store = sessions.NewCookieStore(key)
)

type UserData struct {
	Email    []string `json:"Email"`
	Password []string `json:"Password"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var data User = User{}

	if r.URL.Path != "/testaccueil" {
		http.NotFound(w, r)
		return
	}

	session, _ := store.Get(r, "authenticated")
	auth := session.Values["authenticated"]

	fmt.Println([]byte(auth.(string)))
	fmt.Println(string([]byte(auth.(string))))
	json.Unmarshal([]byte(auth.(string)), &data)
	fmt.Println(data)

	tmpl, _ := template.ParseFiles("./accueil.html")
	tmpl.Execute(w, data)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inscription-connexion" {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template.ParseFiles("./accueil.html")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}

	session, _ := store.Get(r, "authenticated")

	if _, ok := r.PostForm["Submit"]; ok {
		res, _ := json.Marshal(r.PostForm)
		session.Values["authenticated"] = string(res)
		session.Save(r, w)
		http.Redirect(w, r, "/testaccueil", http.StatusFound)
		return
	}

	tmpl.Execute(w, nil)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {

}
