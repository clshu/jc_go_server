package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var layoutPath = TemplatePath{
	Dir:       "views/layouts/",
	Extension: ".gohtml",
}

var templatePath = TemplatePath{
	Dir:       "views/",
	Extension: ".gohtml",
}

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
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

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templatePath.Dir + f + templatePath.Extension
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

type TemplatePath struct {
	Dir       string
	Extension string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}
