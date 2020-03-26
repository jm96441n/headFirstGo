package main

import (
        "log"
        "net/http"
)

func englishHandler(writer http.ResponseWriter, request *http.Request) { 
        message := []byte("Hello Web!")
        _, err := writer.Write(message)
        if err != nil {
                log.Fatal(err)
        }
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) { 
        message := []byte("Salut Web!")
        _, err := writer.Write(message)
        if err != nil {
                log.Fatal(err)
        }
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) { 
        message := []byte("Namaste Web!")
        _, err := writer.Write(message)
        if err != nil {
                log.Fatal(err)
        }
}

func goodbyeHandler(writer http.ResponseWriter, request *http.Request) { 
        message := []byte("Goodbye Web!")
        _, err := writer.Write(message)
        if err != nil {
                log.Fatal(err)
        }
}
func main() {
        http.HandleFunc("/hello", englishHandler)
        http.HandleFunc("/salut", frenchHandler)
        http.HandleFunc("/namaste", hindiHandler)
        http.HandleFunc("/goodbye", goodbyeHandler)
        err := http.ListenAndServe("localhost:8080", nil)
        log.Fatal(err)
}
