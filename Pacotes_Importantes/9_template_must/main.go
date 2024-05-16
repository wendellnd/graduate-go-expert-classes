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
	course := Course{Name: "Go", Workload: 40}
	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}

}
