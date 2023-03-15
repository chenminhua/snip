package main

import (
	"log"
	"net/http"
)

func serve() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}