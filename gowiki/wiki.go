package main

import (
	"fmt"
	"io/ioutil"
)

// Page - a Wiki page
type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2 := load("TestPage")
	fmt.Println(string(p2.Body))
}

func print(p *Page) *Page {
	fmt.Println("I am writing a Page here...")
	return p
}

func load(title string) *Page {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func save() (p *Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}