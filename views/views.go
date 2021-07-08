package views

import "html/template"

func NewView(layout string, files ...string) *View {
	newFiles := append(files,
		"views/layouts/boostrap.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml",
	)

	t, err := template.ParseFiles(newFiles...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
