package views

import (
	"html/template"
	"path/filepath"
)

var layoutPath = LayoutPath{
	Dir:       "views/layouts/",
	Extension: ".gohtml",
}

func NewView(layout string, files ...string) *View {
	newFiles := append(files, layoutFiles()...)

	t, err := template.ParseFiles(newFiles...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutPath.Dir + "*" + layoutPath.Extension)
	if err != nil {
		panic(err)
	}
	return files
}

type View struct {
	Template *template.Template
	Layout   string
}

type LayoutPath struct {
	Dir       string
	Extension string
}
