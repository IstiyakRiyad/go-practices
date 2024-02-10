package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Name		string
	Phone		string
	District	string
	CreatedAt	time.Time
	Position	int64
}

func encodeJSON() {
	d := User{
		Name: "Istiyak Hossain",
		Phone: "01773787127",
		District: "Dinajpur",
		CreatedAt: time.Now(),
		Position: 2,
	}

	j, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(j[:]))
}

func decodeJSON() {
	j := []byte(`
		{
			"Name":"Istiyak Hossain Riyad",
			"Phone":"01773787127",
			"District":"Rajshahi",
			"CreatedAt":"2024-02-10T15:15:10.569764324+06:00",
			"Position": 1
		}
	`)

	var d User
	err := json.Unmarshal(j, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}

func main() {
	encodeJSON()
	decodeJSON()
}



