package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialiser le routeur
	r := mux.NewRouter()
	fmt.Print("Écouté sur le port 8080...\n")

	// Définir une route pour le fichier script.js
	r.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		js, err := ioutil.ReadFile("script.js")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprint(w, string(js))
	})

	// Définir une route pour la page d'accueil
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := ioutil.ReadFile("static/error.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		css, err := ioutil.ReadFile("static/error.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		image, err := ioutil.ReadFile("static/images/404-error-page-examples-best.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<html><head><title>Error 404 page</title><style>%s</style></head><body style=\"background-image: url('data:image/jpg;base64,%s')\">%s</body></html>", string(css), base64.StdEncoding.EncodeToString(image), string(html))
	})

	// Définir une route pour la page de connexion
	r.HandleFunc("/login_page.html", func(w http.ResponseWriter, r *http.Request) {
		html, err := ioutil.ReadFile("login_page.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		css, err := ioutil.ReadFile("static/style.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		image, err := ioutil.ReadFile("./static/images/background.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<html><head><title>Login Page</title><style>%s</style></head><body style=\"background-image: url('data:image/png;base64,%s')\">%s</body></html>", string(css), base64.StdEncoding.EncodeToString(image), string(html))
	})

	// Définir une route pour récupérer les données de la base de données
	r.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var (
			id   int
			name string
		)

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", r))
}
