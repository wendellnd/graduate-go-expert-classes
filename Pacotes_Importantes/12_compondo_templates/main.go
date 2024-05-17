package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	courses := []Course{
		{Name: "Go", Workload: 40},
		{Name: "Node.js", Workload: 50},
		{Name: "React", Workload: 10},
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}

}
