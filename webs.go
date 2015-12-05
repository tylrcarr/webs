package main

import "net/http"

func webs(w http.ResponseWriter, r *http.Request) {
	println("he's so tall and handsome as hell")
	http.ServeFile(w, r, "www/"+r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", webs)
	http.ListenAndServe(":8420", nil)
}
