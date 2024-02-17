package main

import (
	"html/template"
	"os"
)


type Pet struct {
	Name	string
	Sex		string
	Intact	bool
	Age		string
	Breed	string
}

func main() {
	dogs := []Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
		{
			Name:	"Bruce Wayne",
			Sex:	"Male",
			Intact:	false,
			Age:	"3 years, 8 months",
			Breed:	"Chihuahua",
		},
	}

	tmplFile := "pets.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, dogs); err != nil {
		panic(err)
	}
}




