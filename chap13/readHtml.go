package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
)

type urlResponse struct {
        name string
        size int
}

func responseSize(url string, channel chan urlResponse) {
        fmt.Println("Getting", url)
        response, err := http.Get(url)
        if err != nil {
                log.Fatal(err)
        }

        defer response.Body.Close()

        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
                log.Fatal(err)
        }

        resp := urlResponse{name: url, size: len(body)}
        channel <- resp
}

func main() {
        sizes := make(chan urlResponse)
        urls := []string{"https://www.example.com", "https://www.golang.org", "https://www.golang.org/doc"}
        for _, url := range urls {
                go responseSize(url, sizes)
        }

        for i := 0; i < len(urls); i++ {
                resp := <-sizes
                fmt.Printf("Response\n\t%s\n\t%d\n", resp.name, resp.size)
        }
}
