package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
)


type Page struct {
	Title	string
	Body	[]byte
}


func (p *Page) save() error {
	fileName := p.Title + ".txt"

	return os.WriteFile(fileName, p.Body, 0600);
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return &Page{
		Title: title,
		Body: data,
	}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	fileName := tmpl + ".html"

	if err := templates.ExecuteTemplate(w, fileName, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]


	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return;
	}

	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	page, err := loadPage(title)
	if err != nil {
		page = &Page{
			Title: title,
		}
	}
	page.save()

	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
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


func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}


	// page := Page{
	// 	Title: "Page Title",
	// 	Body: []byte("This is the body of the page."),
	// }
	//
	// page.save()
	//
	// pageData, err := loadPage("Page Title he")
	// if(err != nil) {
	// 	fmt.Println("Page not found")
	// }
	//
	// fmt.Println(string(pageData.Body))
}



























