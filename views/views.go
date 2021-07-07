package views

import "html/template"

func NewView(files ...string) *View {
	newFiles := append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(newFiles...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
