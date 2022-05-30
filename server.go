package main

import (
	pckg "Forum/Forum"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type StructTest struct {
	Users  []pckg.User
	Posts  []pckg.Post
	Topics []pckg.Post
}

func main() {
	db := pckg.InitDatabase("forum.db")
	defer db.Close()

	home, err := template.ParseFiles("./pages/accueil.html")
	if err != nil {
		fmt.Println(err)
	}

	logsign, err := template.ParseFiles("./pages/login-signin.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", pckg.HandleHome)

	http.HandleFunc("/inscription-connexion", func(w http.ResponseWriter, r *http.Request) {
		logsign.Execute(w, "")
	})
	http.HandleFunc("/login-signin", pckg.HandleLogin)
	http.HandleFunc("/logout", pckg.HandleLogout)

	http.HandleFunc("/profil", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	http.HandleFunc("/topic", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	http.HandleFunc("/a-propos", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	http.HandleFunc("/get-posts", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	// routes API

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userList := pckg.GetTable(db, "user")
		a := pckg.GetUserRows(userList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		postList := pckg.GetTable(db, "post")
		a := pckg.GetPostRows(postList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	http.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		topicList := pckg.GetTopic(db, "post")
		a := pckg.GetPostRows(topicList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
