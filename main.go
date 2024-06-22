package main

import (
	"go_ProjectSD/handler"
	"net/http"
)

func WebProgram() {
	http.HandleFunc("/login", handler.LoginAdminHandler)
	http.HandleFunc("/register", handler.RegisterAdminHandler)
	http.HandleFunc("/admins", handler.ListAdminsHandler)
	http.HandleFunc("/login-history", handler.LoginHistoryHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.ListenAndServe(":8080", nil)
}

func main() {
	// MainProgram()
	WebProgram()
}
