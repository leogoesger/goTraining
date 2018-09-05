/*
1. Take the previous program in the previous folder and change it so that:
* a template is parsed and served
* you pass data into the template
*/

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
