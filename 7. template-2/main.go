package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)


type Pet struct {
	Name	string
	Sex		string
	Intact	bool
	Age		string
	Breed	[]string
}


func arrayAndOtherFunctions(dogs []Pet) {
	tmplFile := "pets.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("hello world")

	if err := tmpl.Execute(os.Stdout, dogs); err != nil {
		panic(err)
	}
}

func customFunc(dogs []Pet) {
	funcMap := template.FuncMap{
		"dec" : func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
		"join": strings.Join,
	}

	fileName := "lastPet.tmpl"
	tmpl, err := template.New(fileName).Funcs(funcMap).ParseFiles(fileName)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, dogs); err != nil {
		panic(err)
	}
}

func htmlRendering(dogs []Pet) {
	funcMap := template.FuncMap{
		"dec" : func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
		"join": strings.Join,
	}

	fileName := "petsHtml.tmpl"
	tmpl, err := template.New(fileName).Funcs(funcMap).ParseFiles(fileName)
	if err != nil {
		panic(err)
	}

	// Write the html file
	file, err := os.Create("pets.html")
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(file, dogs); err != nil {
		panic(err)
	}
}

func main() {
	dogs := []Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  []string{"German Shepherd", "Pitbull"},
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  []string{"German Shepherd", "Border Collie"},
		},
		{
			Name:	"Bruce Wayne",
			Sex:	"Male",
			Intact:	false,
			Age:	"3 years, 8 months",
			Breed:	[]string{"Chihuahua"},
		},
	}
	
	// arrayAndOtherFunctions(dogs)
	// customFunc(dogs)
	htmlRendering(dogs)
}




