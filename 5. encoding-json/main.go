package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func steamingEncoder() {
	file, err := os.OpenFile("user.json", os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(file)
	
	user := User{"Istiyak Hossain", "01773787127", "Rajshahi", time.Now(), 1,}

	if err := enc.Encode(&user); err != nil {
		log.Fatal(err)
	}
}

func streamingDecoder() {
	// Only Open for reading 
	file, err := os.Open("user.json")
	if err != nil {
		log.Fatal(err)
	}

	// Creating decoder that take input from file
	dec := json.NewDecoder(file)

	var user User
	if err := dec.Decode(&user); err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func streamingStdInOut() {
	dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)

	// Loop continuously
    for {
        var v map[string]interface{}

		// Take input from stdin as json & decode it to v 
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }

		// From all key select only "Name" from the v
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }

		// Show the only "Name" key to the stdout as json
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}

func main() {
	// encodeJSON()
	// decodeJSON()

	// steamingEncoder()
	// streamingDecoder()

	streamingStdInOut()
}



