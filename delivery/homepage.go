package delivery

import (
	"html/template"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, "")
		return
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

		tmpl.Execute(w, nil)
	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}
