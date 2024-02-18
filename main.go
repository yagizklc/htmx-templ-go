package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/yagizklc/htmx-go/templates"
)

func main() {
	component := templates.Hello("John")

	http.Handle("/", templ.Handler(component))
	http.Handle("/click", templ.Handler(templates.ClickResponse("John")))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
