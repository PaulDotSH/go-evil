package handlers

import (
	"errors"
	"fmt"
	settings "go-evil"
	"net/http"
	"os"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println(r.Header)
	case "GET":
		http.ServeFile(w, r, "../frontend/install.html")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": //on post set the cookies and whatever
		fmt.Println(r.Header)
	case "GET":
		http.ServeFile(w, r, "../frontend/login.html")
	}
}

//serve the files as they should be on correct logins
func Panel(w http.ResponseWriter, r *http.Request) {
	//TODO: if this is the first start, make the settings file automatically use the correct endpoint (local ip based)
	if _, err := os.Stat(settings.PwPath); errors.Is(err, os.ErrNotExist) {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
	}
}
