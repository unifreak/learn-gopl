package main

import (
	"text/template"
	"os"
	"log"
	"time"

	"learn/gopl/ch4/github"
)

// template: a string or file containing one or more portions enclosed in
// double braces {{...}}, called action.
//
// action: trigger other behavior. contains an expression in the template language,
// for printing values, selecting struct fields, calling functions and methods,
// expressing control flow such as if-else and range loops, and instantiating ohter templates.

// dot initially refers to the template's parameter.
//
// {{range .Items}}...{{end}} create a loop, dot bound to successive elements of Items.
//
// the | notion makes the result of one operation the argument of another.
const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age: 	{{.CreatedAt | daysAgo}} days
{{end}}`

// Producing output with template is a two-step process
//
// 1. parse the template into a suitable internal representation
// 2. execute it on specific inputs

// .Must helper function makes error handling mor convenient: it accepts a
// template and an error, checks that the error is nil (and panic otherwise),
// and then returns the template.
var report = template.Must(template.New("issuelist"). // .New create and returns a template
	Funcs(template.FuncMap{"daysAgo": daysAgo}). // .FuncMap make `daysAgo` func accessbile within template
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}