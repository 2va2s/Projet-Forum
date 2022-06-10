package main

import (
	pckg "Forum/Forum"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// pckg.Create(db, "category", pckg.Category{}, "Santé", "pink")
	// pckg.Create(db, "category", pckg.Category{}, "Nostalgie", "purple")

	// pqrentPostId, _ := pckg.Create(db, "post", pckg.Post{}, "1 1 1 1 1 1 1 1 1 1 1", 1, "Je suis 1", 1, nil, userId1, "44/44", 0)
	// postId2, _ := pckg.Create(db, "post", pckg.Post{}, "2 2 2 2 ", 0, "Je suis 2", 1, pqrentPostId, userId2, "15/13", 0)
	// pckg.Create(db, "post", pckg.Post{}, "3 3 3 3", 0, "Je suis 3", 1, postId2, userId1, "9312", 0)

	// pqrentPostId2, _ := pckg.Create(db, "post", pckg.Post{}, "11 11 11 11", 1, "Je suis 11", 2, nil, userId1, "25/43", 0)
	// pckg.Create(db, "post", pckg.Post{}, "22 22 22", 0, "Je suis 22", 1, pqrentPostId2, userId2, "35/96", 0)
	// pckg.Create(db, "post", pckg.Post{}, "33 33 33", 0, "Je suis 33", 1, pqrentPostId2, userId1, "14/04", 0)

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
		fmt.Println("aaaaaaaa")
		pckg.HandleSignin(w, r, db)
	})

	rr.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		pckg.HandleLogin(db, w, r)
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

	rr.HandleFunc("/Apropos", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/aproposde.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	rr.HandleFunc("/Cgu", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/cgu.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	rr.HandleFunc("/Support", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/support.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		r.ParseForm()
		// objet := r.Form.Get("objetsupport")
		// corps := r.Form.Get("corpssupport")
		// result := objet + "\n" + corps
		// fmt.Println(result)
		http.Redirect(w, r, "/", http.StatusFound)
		//joindre le ticket a la bdd pour l'afficher dans le profil du superadmin
	})

	rr.HandleFunc("/Equipe", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/equipe.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	rr.HandleFunc("/UpdateVote", func(w http.ResponseWriter, r *http.Request) {
		var params pckg.UpdateVoteParams
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &params)
		// fmt.Println(params)
		w.Write([]byte(pckg.UpdateVotes(db, params.Table, params.Value, params.Field, params.Value2, params.Where)))

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

	rr.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		topicList := pckg.Get(db, "category", "")
		a := pckg.GetCategoryRows(topicList)
		json, _ := json.Marshal(a)
		w.Write(json)
	})

	http.ListenAndServe(":8080", rr)
}
