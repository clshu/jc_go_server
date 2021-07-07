package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Dog   string
	Int   int
	Float float32
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := User{
		Name:  "John Smith",
		Dog:   "Lucy",
		Int:   32,
		Float: 4.23,
		Slice: []string{"Happy", "Life"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
