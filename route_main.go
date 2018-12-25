package main

import (
  "log"
  "net/http"
  "html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
  data := "Hello world from template!"

  templates := template.Must(template.ParseFiles("templates/layout.html"))
  if err := templates.ExecuteTemplate(w, "layout", data); err != nil {
    log.Fatal(err)
  }
}
