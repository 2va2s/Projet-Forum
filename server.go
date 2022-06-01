package main

import (
	pckg "Forum/Forum"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type StructTest struct {
	Users  []pckg.User
	Posts  []pckg.Post
	Topics []pckg.Post
}

func main() {
	rr := mux.NewRouter()
	db := pckg.InitDatabase("forum.db")
	defer db.Close()
	// pckg.CreateUser(db, "kanye", "mdpdezinzin", "", 603504132, "")
	// pckg.CreateUser(db, "booba", "mdpdezinzin", "", 6035041384, "")
	// pckg.CreatePost(db, "Ceci est le topic 1", 1, 1, "")

	home, err := template.ParseFiles("./pages/accueil.html")
	if err != nil {
		fmt.Println(err)
	}

	logsign, err := template.ParseFiles("./pages/login-signin.html")
	if err != nil {
		fmt.Println(err)
	}

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	fileServer := http.FileServer(http.Dir("./static"))
	rr.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))

	// A SUPPRIMER
	rr.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./pages/components/postCard.html")
	})
	// A SUPPRIMER

	rr.HandleFunc("/", pckg.HandleHome)

	rr.HandleFunc("/connexion-inscription", func(w http.ResponseWriter, r *http.Request) {
		logsign.Execute(w, "")
	})

	rr.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		pckg.HandleSignin(w, r, db)
	})

	rr.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		pckg.HandleLogin(w, r, db)
	})

	rr.HandleFunc("/logout", pckg.HandleLogout)

	rr.HandleFunc("/profil/{userID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userID"]
		fmt.Print(userId)
		http.ServeFile(w, r, "./pages/accueil.html")
	})

	rr.HandleFunc("/topic", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	rr.HandleFunc("/a-propos", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	// routes API

	rr.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userList := pckg.GetTable(db, "user")
		a := pckg.GetUserRows(userList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	rr.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		postList := pckg.GetTable(db, "post")
		a := pckg.GetPostRows(postList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	rr.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		topicList := pckg.GetTopic(db, "post")
		a := pckg.GetPostRows(topicList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	http.ListenAndServe(":8080", rr)
}
