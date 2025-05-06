package main

import (
	"fmt"
	"log"
	"net/http"
)

var visitCount int

func handler(w http.ResponseWriter, r *http.Request) {
	visitCount++
	fmt.Fprintln(w, "<h1>Hello, 世界</h1>")
	fmt.Fprintf(w, "<p>Посетителей: %d</p>", visitCount)

	fmt.Fprintf(w, `<img src="/static/image.jpg" alt="Картинка" style="width:500px;">`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)

	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
