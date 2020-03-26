package guestbook

import (
        "bufio"
        "fmt"
        "os"
        "log"
)

type Guestbook struct {
        Lines []string
        Count int 
}

func check(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func AddSignature (signature string, filename string) {
        options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
        file, err := os.OpenFile(filename, options, os.FileMode(0600))
        check(err)
        _, err = fmt.Fprintln(file, signature)
        check(err)
        err = file.Close()
        check(err)
}

func GetGuestbook (filename string) *Guestbook{
        var lines []string
        file, err := os.Open(filename)
        if os.IsNotExist(err) {
                fmt.Println("FILE NOT FOUND")
                return nil
        }

        check(err)
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                lines = append(lines, scanner.Text())
        }

        check(scanner.Err())
        gb := Guestbook{Lines: lines, Count: len(lines)}
        return &gb
}
