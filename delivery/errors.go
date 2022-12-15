package delivery

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Error struct {
	Status           int
	StatusTextandInt string
	MessageError     string
}

func Errors(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error parsing file", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	error1 := Error{status, strconv.Itoa(status) + " " + http.StatusText(status), message}
	if err := t.Execute(w, error1); err != nil {
		log.Print(err)
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error executing file", http.StatusInternalServerError)
		return
	}
}
