package Forum

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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

	if auth != nil {
		json.Unmarshal([]byte(auth.(string)), &data)
	}
	// fmt.Println(data)

	tmpl, _ := template.ParseFiles("./pages/accueil.html")
	tmpl.Execute(w, data)
}

func HandleSignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}
	// checkEmptyField(r)

	if r.URL.Path != "/signin" {
		http.NotFound(w, r)
		return
	}

	Create(db, "user", User{}, r.Form.Get("pseudo"), encrypt(r.Form.Get("password")), r.Form.Get("mail"), r.Form.Get("number"), "")
}

func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// r.Form pour recup valeurs de form
	// db.USer == r.Form
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	///////////////// LOGIN ////////////////

	tmpl, _ := template.ParseFiles("./pages/accueil.html")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}

	// checkLogin(db, r.Form.Get("pseudo"), )

	session, _ := store.Get(r, "cookie-forum")

	if _, ok := r.PostForm["Submit"]; ok {
		fmt.Println("user logged in")
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

func encrypt(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	a := hex.EncodeToString(hasher.Sum(nil))
	return a
}
