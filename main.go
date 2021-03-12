package main

import (
    "fmt"
     "net/http"
 )

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
         fmt.Println("Go Web Hello World!")
     })

     http.ListenAndServe(":31080", nil)
 }
