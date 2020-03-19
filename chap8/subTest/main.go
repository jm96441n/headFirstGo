package main

import (
        "fmt"
        "github.com/headfirstgo/magazine"
)

func main() {
        var s magazine.Subscriber
        s.Rate = 4.99
        s.Name = "Aman Singh"
        s.Active = true
        fmt.Println(s.Name)
        fmt.Println(s.Rate)
        fmt.Println(s.Active)

        s2 := magazine.Subscriber{Name: "Beth Ross", Rate: 4.99, Active: true}

        fmt.Println(s2.Name)
        fmt.Println(s2.Rate)
        fmt.Println(s2.Active)

        employee := magazine.Employee{Name: "Joy Carr", Salary: 60000}
        fmt.Println(employee.Name)
        fmt.Println(employee.Salary)

        address := magazine.Address{Street: "123 Oak St.", City: "Omaha", State: "NE", PostalCode: "68111"}
        employee.HomeAddress = address
        fmt.Println(employee.HomeAddress)

        fmt.Printf("%#v\n", s2.Address)
        s2.Address.Street = "1313 Mockingbird Lane"
        s2.Address.City = "Brooklyn"
        s2.Address.State = "NY"
        s2.Address.PostalCode = "123445"
        fmt.Printf("%#v\n", s2.Address)
}
