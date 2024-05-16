package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("./template.html"))
		courses := []Course{
			{Name: "Go", Workload: 40},
			{Name: "Node.js", Workload: 50},
			{Name: "React", Workload: 10},
		}
		err := t.Execute(w, courses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8282", nil)
}
