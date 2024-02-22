package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)


type Page struct {
	Title	string
	Body	[]byte
}


func (p *Page) save() error {
	fileName := p.Title + ".txt"

	return os.WriteFile("data/" + fileName, p.Body, 0600);
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"

	data, err := os.ReadFile("data/" + fileName)
	if err != nil {
		return nil, err
	}

	return &Page{
		Title: title,
		Body: data,
	}, nil
}

var templates = template.Must(template.ParseFiles("tmpl/view.tmpl", "tmpl/edit.tmpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	fileName := tmpl + ".tmpl"

	if err := templates.ExecuteTemplate(w, fileName, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	// FindStringSubmatch => 0 => full string, 1...n => sub strings () () () # those which are inside () in regex
	m := validPath.FindStringSubmatch(r.URL.Path)

	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}

	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return;
	}

	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{
			Title: title,
		}
	}
	page.save()

	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	page := Page{
		Title: title,
		Body: []byte(body),
	}
	
	if err := page.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, page string)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// FindStringSubmatch => 0 => full string, 1...n => sub strings () () () # those which are inside () in regex
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}
}



























