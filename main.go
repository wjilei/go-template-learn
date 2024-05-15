package main

import (
	"embed"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed views
var views embed.FS

//go:embed static
var static embed.FS

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Errors    map[string]string
}

var contacts []Contact = []Contact{
	{Id: 1, FirstName: "John", LastName: "Smith", Phone: "123-456-7890", Email: "john@example.comz"},
	{Id: 2, FirstName: "Dana", LastName: "Crandith", Phone: "123-456-7890", Email: "dcran@example.com"},
	{Id: 3, FirstName: "Edith", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en@example.com"},
}

type ContactManager struct {
	contacts []Contact
}

var manager *ContactManager = NewContactManager()

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: contacts,
	}
}

func (m *ContactManager) Search(s string) []*Contact {
	var ret []*Contact
	for i := range m.contacts {
		c := m.contacts[i]
		if strings.Contains(c.FirstName, s) {
			ret = append(ret, &c)
		}
	}
	return ret
}

func (m *ContactManager) All() []*Contact {
	var ret []*Contact
	for i := range m.contacts {
		c := m.contacts[i]
		ret = append(ret, &c)
	}
	return ret
}

func (m *ContactManager) Add(c *Contact) error {
	for i := range m.contacts {
		c1 := m.contacts[i]
		if c1.FirstName == c.FirstName {
			if c.Errors == nil {
				c.Errors = make(map[string]string)
			}
			c.Errors["FirstName"] = "already exists"
			// c.Errors["LastName"] = ""
			// c.Errors["Email"] = ""
			// c.Errors["Phone"] = ""

			return errors.New("replicated")
		}
	}
	m.contacts = append(m.contacts, *c)
	return nil
}

func main() {
	t := &Template{
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
