package main

import (
        "guestbook"
        "log"
        "net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
        placeholder := []byte("Guestbook")
        _, err := writer.Write(placeholder)
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
        http.HandleFunc("/guestbook", viewHandler)
        err := http.ListenAndServe("localhost:8080", nil)
        log.Fatal(err)
}
