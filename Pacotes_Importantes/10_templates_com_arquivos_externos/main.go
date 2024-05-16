package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	courses := []Course{
		{Name: "Go", Workload: 40},
		{Name: "Node.js", Workload: 50},
		{Name: "React", Workload: 10},
	}

	t := template.Must(template.New("template.html").ParseFiles("./template.html"))
	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}

}
