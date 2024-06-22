package handler

import (
	"encoding/json"
	"go_ProjectSD/controller"
	"go_ProjectSD/node"
	"html/template"
	"net/http"
	"strconv"
)

// Handler untuk registrasi admin
func RegisterAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data node.DataAdmin
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	controller.CRegisterAdmin(data)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Admin registered successfully"))
}

// Handler untuk login admin
func LoginAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		adminID, err := controller.CLoginAdmin(creds.Username, creds.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Admin logged in successfully. Admin ID: " + strconv.Itoa(adminID)))
	} else {
		http.Error(w, "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
	}
}

// Handler untuk daftar admin
func ListAdminsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	admins := controller.CDaftarAdmin()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admins)
}

// Handler untuk riwayat login admin
func LoginHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	history := controller.CLoginHistory()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
