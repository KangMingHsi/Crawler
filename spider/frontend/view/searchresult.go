package view

import (
	"html/template"
	"io"
	"spider/frontend/model"
)

// SearchResultView ;
type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView ;
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename),
		),
	}
}

// Render ;
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
