package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("---Text Templates---")

	// Define a template with various features
	const tmpl = `
Hello, {{.Name}}!

You have {{if .Age}}{{.Age}}{{else}}an unknown age{{end}}.

Your hobbies:
{{range .Hobbies}}
- {{. | toUpper}}
{{end}}

Custom message:
{{with .CustomMessage}}{{.}}{{else}}No custom message provided{{end}}
`

	// Define a custom function map
	funcMap := template.FuncMap{
		"toUpper": strings.ToUpper, // Custom function to convert string to uppercase
	}

	// Create a template object with functions
	t := template.Must(template.New("example").Funcs(funcMap).Parse(tmpl))

	// Define data
	data := struct {
		Name          string
		Age           int
		Hobbies       []string
		CustomMessage string
	}{
		Name:          "Alice",
		Age:           30,
		Hobbies:       []string{"reading", "hiking", "coding"},
		CustomMessage: "Welcome to Go Templates!",
	}

	// Execute the template with the data
	err := t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
