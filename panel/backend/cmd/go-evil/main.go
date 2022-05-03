package main

import (
	"errors"
	"fmt"
	settings "go-evil"
	"go-evil/parser"
	"net/http"
	"os"
	"path"
)

func init() {

}

//when checking the auth make sure to check the string manually to avoid timming based attacks

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/setup", setup)

	http.HandleFunc("/", panel)

	//serve js and css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../frontend/assets/"))))

	foo()

	http.ListenAndServe(settings.Endpoint, nil)
}

func foo() {
	fmt.Println(parser.Parse(path.Join(settings.RansomwarePath, "settings.go")))
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
	if _, err := os.Stat(settings.PwPath); errors.Is(err, os.ErrNotExist) {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
	}
}

func setup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../frontend/install.html")
}
