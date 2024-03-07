package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() err= $v", err)
		return
	}
	fmt.Fprintf(w, "POST successfully")
	name := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Password= %s\n", password)
}

func main() {
	filerServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", filerServer)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
