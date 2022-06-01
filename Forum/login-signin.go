package Forum

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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

func checkLogin(db *sql.DB, pseudo string, mdp string) {

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

func HandleSignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}
	// checkEmptyField(r)

	if r.URL.Path != "/signin" {
		http.NotFound(w, r)
		return
	}

	// number, _ := strconv.Atoi(r.Form.Get("number"))
	// CreateUser(db, r.Form.Get("pseudo"), encrypt(r.Form.Get("password")), r.Form.Get("mail"), number, "")
	create(db, "user", User{}, r.Form.Get("pseudo"), encrypt(r.Form.Get("password")), r.Form.Get("mail"), r.Form.Get("number"), "")
}

func ParseTable(model interface{}, table string) (a string) {

	result := ""
	e := reflect.ValueOf(model)
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		result += string(varName) + ", "
	}
	result = result[:len(result)-2]
	return result
}

func create(db *sql.DB, table string, model interface{}, t ...string) {

	result := "INSERT INTO " + table + " (" + ParseTable(model, "user") + ")" + " VALUES " + "("
	for i := 0; i < len(t); i++ {
		result += t[i] + ", "
		if i < len(t)-1 {
			result += t[i]
		}
	}
	fmt.Println("1", result)

	result = result[:len(result)-1]
	result += ")"

	// request, _ := db.Exec(result)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(result)
	// return request.LastInsertId()

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

func checkEmptyField(r *http.Request) string {
	// mdp := r.Form.Get("Password")
	// if len(mdp) < 4 {
	// 	return "le champs mot de passe doit contenir au moins 4 caractères"
	// }
	return ""
}

func encrypt(pwd string) string {

	hasher := md5.New()
	hasher.Write([]byte(pwd))
	a := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(a)
	return a
}
