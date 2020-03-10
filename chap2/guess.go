// guess challenges players to guess random numbers
package main

import (
	"fmt"
        "bufio"
        "strings"
        "strconv"
        "os"
        "log"
        "time"
	"math/rand"
)

func main() {
        seconds := time.Now().Unix()
        rand.Seed(seconds)
        target := rand.Intn(100) + 1
        fmt.Println("I've chosen a random number between 1 and 100.")
        fmt.Println("Can you guess it?")

        success := false
        for i := 0; i < 10; i++ {
                fmt.Print("Enter your guess (", 10 - i, " tries remaining): ")
                reader := bufio.NewReader(os.Stdin)
                input, err := reader.ReadString('\n')
                errorCheck(err)
                input = strings.TrimSpace(input)
                guess, err := strconv.Atoi(input)
                errorCheck(err)
                if guess < target {
                        fmt.Println("Oops. Your guess was LOW.")
                } else if guess > target {
                        fmt.Println("Oops. Your guess was HIGH.")
                } else if guess == target {
                        fmt.Println("That's it!!")
                        break
                }
        }
        if !success {
                fmt.Println("You didn't guess it, the number I was thinking of was:", target)
        }
}

func errorCheck (err error) {
        if err != nil {
                log.Fatal(err)
        }
}
