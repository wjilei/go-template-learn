package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed views
var views embed.FS

//go:embed static
var static embed.FS

func main() {
	t := &Template{
		path:      "views",
		templates: template.Must(template.ParseFS(views, "views/*.html")),
		// templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "static",
		Filesystem: http.FS(static),
	}))
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/contacts")
	})
	e.GET("/contacts", func(c echo.Context) error {
		var contacts []*Contact
		search := c.QueryParam("q")
		if search != "" {
			contacts = manager.Search(search)
		} else {
			contacts = manager.All()
		}
		data := struct {
			SearchStr string
			Contacts  []*Contact
		}{
			SearchStr: search,
			Contacts:  contacts,
		}
		return c.Render(200, "index.html", &data)
	})

	e.GET("/contacts/:id", func(c echo.Context) error {
		return nil
	})

	e.GET("/contacts/:id/edit", func(c echo.Context) error {
		return nil
	})

	e.GET("/contacts/new", func(c echo.Context) error {
		return c.Render(http.StatusOK, "new.html", nil)
	})

	e.POST("/contacts/new", func(c echo.Context) error {
		var contact Contact
		contact.FirstName = c.FormValue("first_name")
		contact.Email = c.FormValue("email")
		contact.LastName = c.FormValue("last_name")
		contact.Phone = c.FormValue("phone")

		if err := manager.Add(&contact); err != nil {
			return c.Render(http.StatusOK, "new.html", &contact)
		}
		return c.Redirect(http.StatusFound, "/contacts")
	})

	if err := e.Start("localhost:8080"); err != nil {
		log.Panic(err)
	}
}
