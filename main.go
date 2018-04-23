package main

import (
    "bytes"
    "io"
    "net/http"
)

//go:generate go-bindata -prefix "site/" -pkg main -o bindata.go site/...

func static_handler(rw http.ResponseWriter, req *http.Request) {
  var path string = req.URL.Path
  if path == "" {
    path = "index.html"
  }
  if path == "index.js" {
    path = "index.js"
  }
  if path == "style.css" {
    rw.Header().Set("Content-Type", "text/css")
    path = "style.css"
  }
  if path == "magic8ballQuestion.png" {
    rw.Header().Set("Content-Type", "image/png")
    path = "magic8ballQuestion.png"
  }
  if path == "answerside.png" {
    rw.Header().Set("Content-Type", "image/png")
    path = "answerside.png"
  }
  if bs, err := Asset(path); err != nil {
    rw.WriteHeader(http.StatusNotFound)
  } else {
    var reader = bytes.NewBuffer(bs)
    io.Copy(rw, reader)
  }
}

func main() {
  http.Handle("/", http.StripPrefix("/", http.HandlerFunc(static_handler)))
    http.ListenAndServe(":3000", nil)
}
