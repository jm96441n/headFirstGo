package main

import (
        "fmt"
        "log"
        "github.com/headfirstgo/datafile"
)

func main() {
        file, err := os.Open("data.txt")
        errorCheck(err)

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                fmt.Println(scanner.Text())
        }

        err = file.Close()
        errorCheck(err)

        errorCheck(scanner.Err())
}

func errorCheck (err error) {
        if err != nil {
                log.Fatal(err)
        }
}
