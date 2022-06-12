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
	Pseudo   []string `json:"pseudo"`
	Password []string `json:"password"`
	UserID   []string `json:"user_id"`
}
type UserDataConvert struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
	UserID   string `json:"user_id"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var data UserData = UserData{}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	session, _ := store.Get(r, "cookie-forum")
	auth := session.Values["authenticated"]

	tmpl, _ := template.ParseFiles("./pages/accueil.html", "./templates/menu.html")
	data2 := UserDataConvert{}

	if auth != nil {
		// fmt.Println("-")
		// fmt.Println(auth.(string))
		// fmt.Println("-")
		json.Unmarshal([]byte(auth.(string)), &data)
		data2 = UserDataConvert{data.Pseudo[0], data.Password[0], data.UserID[0]}
	}
	tmpl.Execute(w, data2)
}

func HandleSignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}
	if r.URL.Path != "/signin" {
		http.NotFound(w, r)
		return
	}

	_, err := Create(db, "user", User{}, r.Form.Get("pseudo"), Encrypt(r.Form.Get("password")), r.Form.Get("mail"), r.Form.Get("number"), "")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func HandleLogin(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var dbPseudo string
	var dbPwd string
	var dbId string

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", 500)
	}

	pseudo := r.Form.Get("pseudo")
	password := Encrypt(r.Form.Get("password"))

	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	loginQuery := db.QueryRow("SELECT pseudo, password, id FROM user WHERE pseudo=?", pseudo, password)
	connexion := loginQuery.Scan(&dbPseudo, &dbPwd, &dbId)

	if connexion != nil {
		fmt.Println("error: Wrong password or username. Please try again.")
	} else {

		// tmpl, _ := template.ParseFiles("./pages/accueil.html")

		session, _ := store.Get(r, "cookie-forum")

		r.PostForm["user_id"] = []string{dbId}
		res, _ := json.Marshal(r.PostForm)
		session.Values["authenticated"] = string(res)
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
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

func Encrypt(pwd string) string {
	salt := "Jessica Alba"
	hasher := md5.New()
	hasher.Write([]byte(pwd + salt))
	a := hex.EncodeToString(hasher.Sum(nil))
	return a
}

func IfExists(db *sql.DB, target string, table string, field string) {

	req := "SELECT * FROM " + table + " WHERE " + field + " LIKE " + "'%" + target + "%'"
	res, err := db.Query(req)
	res.Scan(&target, &table, &field)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// fmt.Printf("%v", res)
}

func IsLogged(r *http.Request) {
	var data UserData = UserData{}
	session, _ := store.Get(r, "cookie-forum")
	auth := session.Values["authenticated"]

	json.Unmarshal([]byte(auth.(string)), &data)
	fmt.Println(data)

}

// func checkRegister(db *sql.DB, pseudo string, mail string, number string) bool {
// 	var dbPseudo string
// 	var dbMail string
// 	var dbNumber string
// 	var UserExists bool
// 	// query := db.QueryRow("SELECT pseudo, mail, number FROM user WHERE pseudo=?", pseudo, email, number).Scan(&dbPseudo, &dbEmail, &dbNumber)
// 	query := db.QueryRow("SELECT pseudo, mail, number FROM user WHERE pseudo=?", pseudo, mail, number).Scan(&dbPseudo, &dbMail, &dbNumber)
// 	if dbPseudo == "" {
// 		fmt.Println("user can be created:")
// 		UserExists = false
// 		fmt.Println("ok", query)
// 		return UserExists
// 	} else {
// 		fmt.Println("user already found ! ", dbPseudo, dbNumber, dbNumber)
// 		fmt.Println("Credentials already exists", dbPseudo, dbMail, dbNumber, "please login")
// 		UserExists = true
// 		fmt.Println("nope", query)
// 		return UserExists
// 	}
// }
