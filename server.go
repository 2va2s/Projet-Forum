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

	// test, err := template.ParseFiles("./pages/layout.html", "./templates/menu.html") // route test layout
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// test2, err := template.ParseFiles("./pages/test.html") // route test login/register coulissant
	// if err != nil {
	// 	fmt.Println(err)
	// }

	http.HandleFunc("/", pckg.HandleHome)

	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { // route test layout
	// 	test.Execute(w, r)
	// })

	// http.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) { // route test login/register coulissant
	// 	test2.Execute(w, r)
	// })

	http.HandleFunc("/connexion-inscription", func(w http.ResponseWriter, r *http.Request) {
		logsign.Execute(w, "")
	})

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		pckg.HandleSignin(w, r, db)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		pckg.HandleLogin(w, r, db)
	})

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

	http.HandleFunc("/Apropos", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/aproposde.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	http.HandleFunc("/Cgu", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/cgu.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	http.HandleFunc("/Support", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/support.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		r.ParseForm()
		objet := r.Form.Get("objetsupport")
		corps := r.Form.Get("corpssupport")
		result := objet + "\n" + corps
		fmt.Println(result)
		http.Redirect(w, r, "/", http.StatusFound)
		//joindre le ticket a la bdd pour l'afficher dans le profil du superadmin
	})

	http.HandleFunc("/Equipe", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./pages/equipe.html", "./templates/footer.html", "./templates/logo.html", "./templates/menu.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
