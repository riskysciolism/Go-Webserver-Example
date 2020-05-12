package main

import (
	"Go-Webserver-Example/page"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", page.MakeHandler(page.ViewHandler))
	http.HandleFunc("/edit/", page.MakeHandler(page.EditHandler))
	http.HandleFunc("/save/", page.MakeHandler(page.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

