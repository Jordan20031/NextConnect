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

	db, err := sql.Open("sqlite3", "./basededonnee/data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialiser le routeur
	r := mux.NewRouter()
	fmt.Print("Écouter sur le port 8080...\n")

	// Définir une route pour le fichier script.js
	r.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier script.js
		js, err := ioutil.ReadFile("script.js")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprint(w, string(js))
	})

	// Définir une route pour le fichier changement_pdp.js
	r.HandleFunc("/changement_pdp.js", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier changement_pdp.js
		js, err := ioutil.ReadFile("changement_pdp.js")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprint(w, string(js))
	})

	// Définir une route pour la page d'accueil ("/")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Lire le contenu du fichier login_page.html
			html, err := ioutil.ReadFile("./templates/html/login_page.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Lire le contenu du fichier style.css
			css, err := ioutil.ReadFile("templates/css/style.css")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Lire le contenu du fichier image (background.jpg)
			image, err := ioutil.ReadFile("./static/images/background.jpg")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Définir les en-têtes de réponse pour le type de contenu
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			// Afficher la page HTML avec le CSS et l'image incorporés
			fmt.Fprintf(w, "<html><head><title>Login Page</title><style>%s</style></head><body style=\"background-image: url('data:image/png;base64,%s')\">%s</body></html>", string(css), base64.StdEncoding.EncodeToString(image), string(html))
		} else if r.Method == http.MethodPost {
			// Gérer la soumission du formulaire POST (non implémenté dans le code fourni)
			file, _, err := r.FormFile("image")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			fmt.Fprint(w, "Photo de profil mise à jour avec succès")
		}
	})

	// Définir une route pour la page d'erreur 404 ("/templates/html/login_page.html")
	r.HandleFunc("/templates/html/login_page.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier error404.html
		html, err := ioutil.ReadFile("./templates/html/error404.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier error404.css
		css, err := ioutil.ReadFile("./templates/css/error404.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier image (404-error-page-examples-best.jpg)
		image, err := ioutil.ReadFile("static/images/404-error-page-examples-best.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page d'erreur 404 avec le CSS et l'image incorporés
		fmt.Fprintf(w, "<html><head><title>Error 404 page</title><style>%s</style></head><body style=\"background: url('data:image/jpg;base64,%s') no-repeat center center fixed; background-size: cover;\">%s</body></html>", string(css), base64.StdEncoding.EncodeToString(image), string(html))
	})

	// Définir une route pour la page mainpage.html
	r.HandleFunc("/mainpage.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier mainpage.html
		html, err := ioutil.ReadFile("./templates/html/mainpage.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier mainpage.css
		css, err := ioutil.ReadFile("./templates/css/mainpage.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page choicesubject.html avec le CSS incorporé
		fmt.Fprintf(w, "<html><head><title>Choose Subject</title><style>%s</style></head><body>%s</body></html>", string(css), string(html))
	})

	// Définir une route pour la page createpost.html
	r.HandleFunc("/createpost.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier mainpage.html
		html, err := ioutil.ReadFile("./templates/html/createpost.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier mainpage.css
		css, err := ioutil.ReadFile("./templates/css/createpost.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page createpost.html avec le CSS incorporé
		fmt.Fprintf(w, "<html><head><title>create post</title><style>%s</style></head><body>%s</body></html>", string(css), string(html))
	})

	// Définir une route pour la page contact.html
	r.HandleFunc("/contact.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier contact.html
		html, err := ioutil.ReadFile("./templates/html/contact.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier contact.css
		css, err := ioutil.ReadFile("./templates/css/contact.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page choicesubject.html avec le CSS incorporé
		fmt.Fprintf(w, "<html><head><title>Choose Subject</title><style>%s</style></head><body>%s</body></html>", string(css), string(html))
	})

	// Définir une route pour la page choicesubject.html
	r.HandleFunc("/choicesubject.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier choicesubject.html
		html, err := ioutil.ReadFile("./templates/html/choicesubject.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier choicesubject.css
		css, err := ioutil.ReadFile("./templates/css/choicesubject.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page choicesubject.html avec le CSS incorporé
		fmt.Fprintf(w, "<html><head><title>Choose Subject</title><style>%s</style></head><body>%s</body></html>", string(css), string(html))
	})

	// Définir une route pour la page discussion.html
	r.HandleFunc("/discussion.html", func(w http.ResponseWriter, r *http.Request) {
		// Lire le contenu du fichier discussion.html
		html, err := ioutil.ReadFile("./templates/html/discussion.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lire le contenu du fichier discussion.css
		css, err := ioutil.ReadFile("./templates/css/discussion.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définir les en-têtes de réponse pour le type de contenu
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Afficher la page discussion.html avec le CSS incorporé
		fmt.Fprintf(w, "<html><head><title>Discussion</title><style>%s</style></head><body>%s</body></html>", string(css), string(html))
	})

	// Définir une route pour récupérer les données de la base de données ("/data")

	r.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Exécuter une requête SQL pour récupérer les données de la table "users"
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

		// Parcourir les lignes de résultats et afficher les données
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

	// Définir une route pour servir les fichiers CSS statiques
	fs := http.FileServer(http.Dir("./templates/css/"))
	r.PathPrefix("/templates/css/").Handler(http.StripPrefix("/templates/css/", fs))

	fmt.Printf("Server is running on http://localhost:8080/\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

/** example of how to link html to db to store a newly created discution.
func handalerSendDiscution(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./basededonnee/data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// (vous devrez ajouter la logique pour récupérer l'utilisateur et ses informations de session)
	user := getUserFromSession(r)
	if user == nil {
		http.Error(w, "Utilisateur non connecté", http.StatusUnauthorized)
		return
	}

	// Lire les données de la requête POST
	err = r.ParseMultipartForm(10 << 20) // 10 MB de taille maximale pour le fichier
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	image, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier image", http.StatusBadRequest)
		return
	}
	defer image.Close()

	titre := r.FormValue("titre")
	description := r.FormValue("description")

	// Appeler la fonction createDiscussion
	err = basededonnee.createDiscussion(db, *user, readImageBytes(image), titre, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Répondre avec un message de succès
	fmt.Fprint(w, "Discussion créée avec succès")
}

// Fonction utilitaire pour lire les octets d'un fichier
func readImageBytes(file multipart.File) []byte {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}
*/
