package main

import (
	"SIMPLEWEBAPP/API"
	"SIMPLEWEBAPP/DB"
	"fmt"
	"net/http"
)

func main() {
	err := DB.OpenDB()
	if err != nil {
		fmt.Println(err)
	}

	API.LoadTemplates()

	http.HandleFunc("/", API.RootHandler)
	http.HandleFunc("/login", API.LoginHandler)
	http.HandleFunc("/register", API.RegisterHandler)
	http.HandleFunc("/authenticateR", API.AuthenticateRegister)
	http.HandleFunc("/authenticate", API.AuthenticateHandler)
	http.HandleFunc("/dashboard", API.DashboardHandler)
	http.ListenAndServe(":8080", nil)
}
