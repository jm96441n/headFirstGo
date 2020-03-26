package main

import (
        "guestbook"
        "html/template"
        "log"
        "net/http"
)

func filePath() string {
        return "../src/guestbookWeb/signatures.txt"
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
        gb := guestbook.GetGuestbook(filePath())
        html, err := template.ParseFiles("../src/guestbookWeb/view.html")
        check(err)
        err = html.Execute(writer, gb)
        check(err) 
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
        html, err := template.ParseFiles("../src/guestbookWeb/new.html")
        check(err)
        err = html.Execute(writer, nil)
        check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
        signature := request.FormValue("signature")
        guestbook.AddSignature(signature, filePath())
        http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func check(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
        http.HandleFunc("/guestbook", viewHandler)
        http.HandleFunc("/guestbook/new", newHandler)
        http.HandleFunc("/guestbook/create", createHandler)
        err := http.ListenAndServe("localhost:8080", nil)
        log.Fatal(err)
}
