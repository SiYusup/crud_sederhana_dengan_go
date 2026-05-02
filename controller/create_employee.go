package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]
			
			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES (?,?,?)", name, npwp, address)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "../employee", http.StatusSeeOther)
			return
		} else if r.Method == http.MethodGet {
			createPath := filepath.Join("views", "create.html")
			template, error := template.ParseFiles(createPath)

			if error != nil {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}

			error = template.Execute(w, nil)

			if error != nil {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}
