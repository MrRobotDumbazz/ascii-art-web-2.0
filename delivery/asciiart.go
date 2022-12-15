package delivery

import (
	"ascii-art-web/internal/service"
	"errors"
	"html/template"
	"net/http"
)

var ErrBadTyping = errors.New("Bad typing")

func asciiart(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		Errors(w, http.StatusNotFound, "")
		return
	}
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		input, ok := r.Form["input"]
		if !ok {
			Errors(w, http.StatusBadRequest, ErrBadTyping.Error())
			return
		}
		fonts, ok := r.Form["fonts"]
		if !ok {
			Errors(w, http.StatusBadRequest, ErrBadTyping.Error())
			return
		}
		if err := service.IsValidInput(input[0]); err != nil {
			Errors(w, http.StatusBadRequest, err.Error())
			return
		}
		ascii, err := service.RunAscii(input[0], fonts[0])
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		t.Execute(w, ascii)
	}
}
