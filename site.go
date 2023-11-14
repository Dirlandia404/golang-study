package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type Page struct {
    Title string
    Body string
}

func main() {
    fmt.Println("Running the site...")

    tmpl := template.Must(template.ParseFiles("index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        p := Page{
            Title: "My Page", // Changed title to English
            Body: "Hello, world!", // Changed body to English
        }

        tmpl.Execute(w, p)
    })

    http.ListenAndServe(":8080", nil)
}
