package backend

import (
	"html/template"
	"net/http"
)

func init() {
	http.Handle("/", MustNew())
}

func MustNew() http.Handler {
	h, err := New()
	if err != nil {
		panic(err)
	}
	return h
}

func New() (h http.Handler, err error) {
	t, err := template.ParseFiles("app/index.html")
	if err != nil {
		return
	}

	h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "index", &struct {
			Title string
		}{
			"{{ name }}",
		})
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
	})

	return
}
