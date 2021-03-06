package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}

type User struct {
	Name  string
	Dog   Dog
	Slice []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := User{
		Name: "John Smith",
		Dog: Dog{
			Name: "Lucy",
		},
		Slice: []string{"Happy", "Life"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
