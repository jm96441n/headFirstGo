package main

import (
        "log"
        "os"
        "text/template"
)

type Part struct {
        Name string
        Count int
}

type Subscriber struct {
        Name string
        Rate float64
        Active bool
}

func check(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func executeTemplate(text string, data interface{}) {
        tmpl, err := template.New("test").Parse(text)
        check(err)
        err = tmpl.Execute(os.Stdout, data)
        check(err)
}

func main() {
        templateText := "Before {{.}}\n{{range .}}In Loop: {{.}}\n{{end}}After Loop: {{.}}\n\n"
        executeTemplate(templateText, []string{"do", "re", "mi"})
        executeTemplate(templateText, []float64{5.10, 2.67, 500.1245})
        executeTemplate(templateText, []float64{})
        executeTemplate(templateText, nil)


        templateText = "Name: {{.Name}}\nCount: {{.Count}}\n\n"
        executeTemplate(templateText, Part{Name: "Fuses", Count: 50})
        executeTemplate(templateText, Part{Name: "Hex Bolts", Count: 25})


        templateText = "Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}{{end}}\n"
        executeTemplate(templateText, Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true})
        executeTemplate(templateText, Subscriber{Name: "Joy Carr", Rate: 4.99, Active: false})
}
