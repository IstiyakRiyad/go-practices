package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func openAndWrite(path string, data string) {
	file, err := os.OpenFile(path, os.O_CREATE| os.O_APPEND | os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// We can use file.Write([]byte) function for byte data
	n, err := file.WriteString(data);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

func pathLibChecking() {
	// Directory from where we running the program
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current working directory: ", wd)
	fmt.Println(path.Base("/hi/hello/hi.txt"))
	fmt.Println(path.Dir("/hi/hello/hi.txt"))
	fmt.Println(path.Ext("/hi/hello/hi.txt"))
	fmt.Println(path.Join("hello", "there", "hi.ts"))
	fmt.Println(path.Split("/hi/hello/hi.txt"))
	fmt.Println(path.IsAbs("hi/hello")) // If he path is given from the root folder 

}

func makeDirectoryAndDirectFileWriteAndDelete(directory, fileName, data string) {
	if err := os.Mkdir(directory, 0777); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	if err := os.WriteFile(directory+ "/" + fileName, []byte(data), 0666); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(directory); err != nil {
		log.Fatal(err)
	}
}

func pipeFunction() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "Writing to the the writer pile\n")
		w.Close()
	}()

	// Copy
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func readFromFile() {
	file, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
}

func openAndReadFromFile() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 50)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main() {
	openAndWrite("text.txt", "Hello World")

	readFromFile()
	openAndReadFromFile()
	
	// makeDirectoryAndDirectFileWriteAndDelete("files", "secret.txt", "Secret data")


	// pathLibChecking()


	// pipeFunction()

}





