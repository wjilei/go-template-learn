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
	{Id: 3, FirstName: "Edith1", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en1@example.com"},
	{Id: 4, FirstName: "Edith2", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en2@example.com"},
	{Id: 5, FirstName: "Edith3", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en3@example.com"},
	{Id: 6, FirstName: "Edith4", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en4@example.com"},
	{Id: 7, FirstName: "Edith5", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en5@example.com"},
	{Id: 8, FirstName: "Edith6", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en6@example.com"},
	{Id: 9, FirstName: "Edith7", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en7@example.com"},
	{Id: 10, FirstName: "Edith8", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en8@example.com"},
	{Id: 11, FirstName: "Edith9", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en9@example.com"},
	{Id: 12, FirstName: "Edith10", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en10@example.com"},
	{Id: 13, FirstName: "Edith10", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en11@example.com"},
	{Id: 14, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en12@example.com"},
	{Id: 15, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en13@example.com"},
	{Id: 16, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en14@example.com"},
	{Id: 17, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en15@example.com"},
	{Id: 18, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en16@example.com"},
	{Id: 19, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en17@example.com"},
	{Id: 20, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en18@example.com"},
	{Id: 21, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en19@example.com"},
	{Id: 22, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en20@example.com"},
	{Id: 23, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en21@example.com"},
	{Id: 24, FirstName: "Edith11", LastName: "Neutvaar", Phone: "123-456-7890", Email: "en22@example.com"},
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

func (m *ContactManager) Search(s string, page, limit int) ([]*Contact, int) {
	var ret []*Contact
	for i := range m.contacts {
		c := m.contacts[i]
		if strings.Contains(c.FirstName, s) {
			ret = append(ret, &c)
		}
	}

	total := len(ret)
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	return ret[start:end], total
}

func (m *ContactManager) All(page, limit int) ([]*Contact, int) {
	var ret []*Contact
	for i := range m.contacts {
		c := m.contacts[i]
		ret = append(ret, &c)
	}

	total := len(ret)
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	return ret[start:end], total
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
		if c.Email == c1.Email {
			c.Errors["Email"] = "Email已存在"
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
	if c.Email == "" {
		c.Errors["Email"] = "邮箱不能为空"
		return errors.New("empty email")
	}
	if c.FirstName == "" {
		c.Errors["FirstName"] = "FirstName不能为空"
		return errors.New("empty first name")
	}
	if c.LastName == "" {
		c.Errors["LastName"] = "LastName不能为空"
		return errors.New("empty last name")
	}

	if !regexp.MustCompile(`^.*@.*\..*$`).MatchString(c.Email) {
		c.Errors["Email"] = "邮箱格式错误"
		return errors.New("invalid email format")
	}

	for i := range m.contacts {
		cdb := m.contacts[i]

		if c.Id != cdb.Id {
			if c.Email == cdb.Email {
				c.Errors["Email"] = "邮箱已存在"
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
