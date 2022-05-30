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
	var data UserData = UserData{}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	session, _ := store.Get(r, "cookie-forum")
	auth := session.Values["authenticated"]

	// fmt.Println([]byte(auth.(string)))
	// fmt.Println(string([]byte(auth.(string))))
	if auth != nil {
		json.Unmarshal([]byte(auth.(string)), &data)
	}
	// fmt.Println(data)

	tmpl, _ := template.ParseFiles("./pages/accueil.html")
	tmpl.Execute(w, data)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login-signin" {
		http.NotFound(w, r)
		return
	}

	tmpl, _ := template.ParseFiles("./pages/accueil.html")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}

	session, _ := store.Get(r, "cookie-forum")

	if _, ok := r.PostForm["Submit"]; ok {
		fmt.Println("in")
		res, _ := json.Marshal(r.PostForm)
		session.Values["authenticated"] = string(res)
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	tmpl.Execute(w, nil)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.NotFound(w, r)
		return
	}

	session, _ := store.Get(r, "cookie-forum")
	session.Values["authenticated"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
