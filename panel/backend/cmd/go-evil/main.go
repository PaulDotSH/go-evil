package main

import (
	"errors"
	"net/http"
	"os"
	"path"
)

//TODO: think of a way to make the settings

const dataDir = "./"
const endpoint = "localhost:6031"

var pwPath = path.Join(dataDir, "login")

func init() {

}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/setup", setup)

	http.HandleFunc("/", panel)

	//serve js and css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../frontend/assets/"))))

	http.ListenAndServe(endpoint, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": //on post set the cookies and whatever
	case "GET":
		http.ServeFile(w, r, "../frontend/login.html")
	}
}

//serve the files as they should be on correct logins
func panel(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(pwPath); errors.Is(err, os.ErrNotExist) {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
	}
}

func setup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../frontend/install.html")
}
