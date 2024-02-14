package main

import (
	"html/template"
	"log"
	"os"
)

type Student struct {
	Name	string
	Roll	int
}

type UserItem struct {
	Name	string
	Items	[]string
}

func parseFromString() {
	templateContent := `
		<h1>My Name is {{.Name}}</h1>
		<h2>My Roll is {{.Roll}}</h2>
	`

	tmp := template.New("parseFromString")

	// Parse from string 
	tmp, err := tmp.Parse(templateContent)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the the template and print to stdout
	if err := tmp.Execute(os.Stdout, Student{Name: "<Istiyak>", Roll: 1803094,}); err != nil {
		log.Fatal(err)
	}
}

func parseFromFile() {
	templateFileName:= "parseFromFile.tmpl"

	tmp := template.New(templateFileName)	// here the name of the template and he parsefile name must have to be same

	// Parse from string 
	tmp, err := tmp.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the the template and print to stdout
	if err := tmp.Execute(os.Stdout, Student{Name: "<Istiyak>", Roll: 1803094,}); err != nil {
		log.Fatal(err)
	}
}

func parseFromFileComplexTemplate() {
	templateFileName:= "userItems.tmpl"

	tmp := template.New(templateFileName)	// here the name of the template and he parsefile name must have to be same

	// Parse from string 
	tmp, err := tmp.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	userInfo := UserItem{
		Name: "Istiyak Hossain",
		Items: []string{
			"item 1", 
			"item 2", 
			"item 3",
		},
	}

	// Execute the the template and print to stdout
	if err := tmp.Execute(os.Stdout, userInfo); err != nil {
		log.Fatal(err)
	}
}


func main() {
	// parseFromString()
	// parseFromFile()
	parseFromFileComplexTemplate()
}





