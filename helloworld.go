package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  httpServer := http.NewServeMux()
  httpServer.HandleFunc("/", sayHello)
  port := "8080"
  fmt.Printf("Starting server on port %s ...", port)
  error := http.ListenAndServe(":"+port, httpServer)
  log.Fatal(error)
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintln(writer,"Hello world from go!")
}
