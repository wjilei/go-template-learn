package main

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	fs        embed.FS
	path      string
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t1 := template.Must(t.templates.Clone())
	t1 = template.Must(t1.ParseFS(t.fs, t.path+"/"+name))
	return t1.ExecuteTemplate(w, name, data)
}
