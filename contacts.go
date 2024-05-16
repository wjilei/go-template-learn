package main

import (
	"errors"
	"regexp"
	"strings"
)

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
	if c.Errors == nil {
		c.Errors = make(map[string]string)
	}
	var idMax int = -1
	for i := range m.contacts {
		c1 := m.contacts[i]
		if c1.Id > idMax {
			idMax = c1.Id
		}
		if c1.FirstName == c.FirstName {
			c.Errors["FirstName"] = "already exists"
			return errors.New("replicated")
		}
		if c.LastName == c1.LastName {
			c.Errors["LastName"] = "already exists"
			return errors.New("replicated")
		}
		if c.Email == c1.Email {
			c.Errors["Email"] = "already exists"
			return errors.New("replicated")
		}
		if c.Phone == c1.Phone {
			c.Errors["Email"] = "already exists"
			return errors.New("replicated")
		}
	}
	c.Id = idMax + 1
	m.contacts = append(m.contacts, *c)
	return nil
}

func (m *ContactManager) Update(c *Contact) error {
	pos := -1
	if c.Errors == nil {
		c.Errors = make(map[string]string)
	}
	if c.FirstName == "" {
		c.Errors["FirstName"] = "cannot be empty"
		return errors.New("empty first name")
	}
	if c.LastName == "" {
		c.Errors["LastName"] = "cannot be empty"
		return errors.New("empty last name")
	}
	if c.Email == "" {
		c.Errors["Email"] = "cannot be empty"
		return errors.New("empty email")
	}
	if c.Phone == "" {
		c.Errors["Phone"] = "cannot be empty"
		return errors.New("empty phone")
	}

	if !regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`).MatchString(c.Phone) {
		c.Errors["Phone"] = "invalid phone format"
		return errors.New("invalid phone format")
	}
	if !regexp.MustCompile(`^.*@.*\..*$`).MatchString(c.Email) {
		c.Errors["Email"] = "invalid email format"
		return errors.New("invalid email format")
	}

	for i := range m.contacts {
		cdb := m.contacts[i]

		if c.Id != cdb.Id {
			if c.FirstName == cdb.FirstName {
				c.Errors["FirstName"] = "already exists"
				return errors.New("replicated")
			}
			if c.LastName == cdb.LastName {
				c.Errors["LastName"] = "already exists"
				return errors.New("replicated")
			}
			if c.Email == cdb.Email {
				c.Errors["Email"] = "already exists"
				return errors.New("replicated")
			}
			if c.Phone == cdb.Phone {
				c.Errors["Email"] = "already exists"
				return errors.New("replicated")
			}
		}
		if c.Id == cdb.Id {
			pos = i
		}

		if !regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`).MatchString(c.Phone) {
			c.Errors["Phone"] = "invalid phone format"
			return errors.New("invalid phone format")
		}
		if !regexp.MustCompile(`^.*@.*\..*$`).MatchString(c.Email) {
			c.Errors["Email"] = "invalid email format"
			return errors.New("invalid email format")
		}
	}
	if pos == -1 {
		return errors.New("fatal")
	}
	m.contacts[pos] = *c
	return nil
}

func (m *ContactManager) Get(id int) (*Contact, error) {
	for i := range m.contacts {
		c := m.contacts[i]
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *ContactManager) Delete(id int) {
	var pos int = -1
	for i, c := range m.contacts {
		if c.Id == id {
			pos = i
		}
	}
	if pos == -1 {
		return
	}

	v := m.contacts[:pos]
	v = append(v, m.contacts[pos+1:]...)
	m.contacts = v
}
