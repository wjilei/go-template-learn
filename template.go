package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	path      string
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t1 := template.Must(t.templates.Clone())
	t1 = template.Must(t1.ParseGlob(t.path + "/" + name))
	return t1.ExecuteTemplate(w, name, data)
}
