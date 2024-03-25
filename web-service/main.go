package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}
