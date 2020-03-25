package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "path/filepath"
)

func readDirectory(dirName string) error {
        fmt.Println(dirName)
        files, err := ioutil.ReadDir(dirName)
        if err != nil {
                log.Fatal(err)
        }

        for _, file := range files {
                filePath := filepath.Join(dirName, file.Name())
                if file.IsDir() {
                        err := readDirectory(filePath)
                        if err != nil {
                                return err
                        }
                } else {
                        fmt.Println(filePath)
                }
        }
        return nil
}

func main() {
        err := readDirectory("my_directory") 
        if err != nil {
                log.Fatal(err)
        }
}
