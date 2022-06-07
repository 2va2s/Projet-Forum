package main

import (
	pckg "Forum/Forum"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	rr := mux.NewRouter()
	db := pckg.InitDatabase("forum.db")
	defer db.Close()

	// PAS SUPPRIMER: DECOMMENTER POUR GENERER TABLES EXEMPLE //

	// userId1, _ := pckg.Create(db, "user", pckg.User{}, "akhy deter", "mdp", "aeze@gmail.com", "6314134235235", "")
	// userId2, _ := pckg.Create(db, "user", pckg.User{}, "fifi grognon", "mdp", "aeze@gmail.com", "6314134235235", "")

	// pckg.Create(db, "category", pckg.Category{}, "Nostalgie", "purple")
	// pqrentPostId, _ := pckg.Create(db, "post", pckg.Post{}, "1 1 1 1 1 1 1 1 1 1 1", 1, "Je suis 1", 1, nil, userId1, "44/44", 0)
	// postId2, _ := pckg.Create(db, "post", pckg.Post{}, "2 2 2 2 ", 0, "Je suis 2", 1, pqrentPostId, userId2, "15/13", 0)
	// pckg.Create(db, "post", pckg.Post{}, "3 3 3 3", 0, "Je suis 3", 1, postId2, userId1, "9312", 0)

	// pqrentPostId2, _ := pckg.Create(db, "post", pckg.Post{}, "11 11 11 11", 1, "Je suis 11", 1, nil, userId1, "25/43", 0)
	// pckg.Create(db, "post", pckg.Post{}, "22 22 22", 0, "Je suis 22", 1, pqrentPostId2, userId2, "35/96", 0)
	// pckg.Create(db, "post", pckg.Post{}, "33 33 33", 0, "Je suis 33", 1, pqrentPostId2, userId1, "14/04", 0)

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

	rr.HandleFunc("/post/{postID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["postID"]
		fmt.Print(userId)
		http.ServeFile(w, r, "./pages/topic.html")
	})

	rr.HandleFunc("/a-propos", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, "")
	})

	// routes API

	rr.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userList := pckg.Get(db, "user", "user")
		a := pckg.GetUserRows(userList)
		for i := 0; i < len(a); i++ {
			rr.HandleFunc("/user/"+strconv.Itoa(a[i].ID), func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "./pages/user.html")
			})
		}
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	rr.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		postList := pckg.Get(db, "post", "child")
		a := pckg.GetPostRows(postList)
		for i := 0; i < len(a); i++ {
			rr.HandleFunc("/topic/"+strconv.Itoa(a[i].ID), func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "./pages/topic.html")
			})
		}
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	rr.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		topicList := pckg.Get(db, "post", "topic")
		a := pckg.GetPostRows(topicList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	http.ListenAndServe(":8080", rr)
}
