package main

import (
	"html/template"
	"os"
	"log"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	// We can suppress html/template's auto-escaping behavior, using
	// the named string type template.HTML instead of string.
	var data struct {
		A string 		// untursted plain text -> auto-escape
		B template.HTML // trusted HTML -> supress auto-escape
	}
	data.A = "<b>Hello!</b>" // output: &lt;b&gt;Hello!&lt;/b&gt;
	data.B = "<b>Hello!</b>" // output: <b>Hello!</b>
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}