package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			id := r.Form["id"][0]
			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]

			_, err := db.Exec("UPDATE employee SET name=?, npwp=?, address=? WHERE id=?", name, npwp, address, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "../employee", http.StatusSeeOther)
			return
		} else if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)
			if row.Err() != nil {
				http.Error(w, row.Err().Error(), http.StatusInternalServerError)
				return
			}

			var employee Employee
			err := row.Scan(
				&employee.Id,
				&employee.Name,
				&employee.Npwp,
				&employee.Address,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			editPath := filepath.Join("views", "edit.html")
			template, error := template.ParseFiles(editPath)

			if error != nil {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}

			error = template.Execute(w, data)

			if error != nil {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}
