package main

import (
  "fmt"
  "errors"
  "os"
  "net/http"
  "io"
)

func getRoot(w http.ResponseWriter, r *http.Request){
  fmt.Printf("got / request\n")
  io.WriteString(w, "This is my website\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
  fmt.Printf("got /hello request\n")
  io.WriteString(w, "Hello, HTTP!\n")
}

func main()  {
  mux := http.NewServeMux()
  mux.HandleFunc("/", getRoot)
  mux.HandleFunc("/hello", getHello)

  err := http.ListenAndServe(":3333", mux)

  if errors.Is(err, http.ErrServerClosed) {
    fmt.Printf("server closed\n")
  }else if err != nil {
    fmt.Printf("error starting server: %s\n", err)
    os.Exit(1)
  }  
}
