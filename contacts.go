package main

import (
	"errors"
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

func (m *ContactManager) Update(c *Contact) error {
	pos := -1
	for i := range m.contacts {
		cdb := m.contacts[i]

		if c.Id != cdb.Id && c.FirstName == cdb.FirstName {
			c.Errors["FirstName"] = "already exists"
		}
		if c.Id == cdb.Id {
			pos = i
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
