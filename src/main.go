package main

import (
  "fmt"
  "errors"
  "context"
  "net"
  "net/http"
  "io"
)

const KeyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request){
  ctx := r.Context()

  fmt.Printf("%s:got / request\n", ctx.Value(KeyServerAddr))
  io.WriteString(w, "This is my website\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
  ctx := r.Context()
 
  fmt.Printf("%s:got /hello request\n", ctx.Value(KeyServerAddr))
  io.WriteString(w, "Hello, HTTP!\n")
}

func main()  {
  mux := http.NewServeMux()
  mux.HandleFunc("/", getRoot)
  mux.HandleFunc("/hello", getHello)
  
  ctx, cancelCtx := context.WithCancel(context.Background())
  
  serveOne := &http.Server{
    Addr: ":3333",
    Handler: mux,
    BaseContext: func (l net.Listener) context.Context {
      ctx = context.WithValue(ctx, KeyServerAddr, l.Addr().String())
      return ctx
    },
  } 
  
  serveTwo := &http.Server{
    Addr: ":4444",
    Handler: mux,
    BaseContext: func (l net.Listener) context.Context {
      ctx = context.WithValue(ctx, KeyServerAddr, l.Addr().String())
      return ctx
    },
  }

  go func(){
    err := serveOne.ListenAndServe()
    if errors.Is(err, http.ErrServerClosed){
      fmt.Printf("server one closed\n")
    }else if err != nil {
      fmt.Printf("error listening for server one: %s\n", err)
    }
    cancelCtx()
  }()

  go func(){
    err := serveTwo.ListenAndServe()
    if errors.Is(err, http.ErrServerClosed){
      fmt.Printf("server one closed\n")
    }else if err != nil {
      fmt.Printf("error listening for server one: %s\n", err)
    }
    cancelCtx()
  }()

  <-ctx.Done()
}
