package delivery

import "net/http"

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/ascii-art", asciiart)
	return mux
}
