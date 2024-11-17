package API

import (
	db "SIMPLEWEBAPP/DB"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var templates *template.Template

type user struct {
	ID        int
	Username  string
	Password  string
	CreatedAt string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, templates, "root.tmpl", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, templates, "login.tmpl", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, templates, "register.tmpl", nil)
}

func AuthenticateRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "problem authenticating", http.StatusBadRequest)
	}

	createdAt := time.Now()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
	_, err := db.DB.Exec(query, username, password, createdAt)
	if err != nil {
		http.Error(w, "Failed to register", http.StatusBadRequest)
		fmt.Println(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "problem authenticating", http.StatusBadRequest)
	}

	userChecked := user{}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	query := "SELECT * FROM USERS WHERE username=? and password=?"
	err := db.DB.QueryRow(query, username, password).Scan(&userChecked.ID, &userChecked.Username, &userChecked.Password, &userChecked.CreatedAt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "problem authenticating", http.StatusBadRequest)
	}

	if userChecked.Username == username && userChecked.Password == password {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Error(w, "Failed to login", http.StatusBadRequest)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, templates, "dashboard.tmpl", nil)
}

func RenderPage(w http.ResponseWriter, tmpls *template.Template, name string, data interface{}) {
	err := tmpls.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.tmpl"))
}
