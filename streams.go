package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ioCopy() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func doubleRead() {

	type Person struct {
		Name string
		Age  int
	}
	r := strings.NewReader(`{"Name": "Jack", "Age": 28}`)
	var person, personAgain Person

	_ = json.NewDecoder(r).Decode(&person)
	fmt.Println(person)

	// reset seek pos to start, to read again
	newSeekPos, _ := r.Seek(0, io.SeekStart)
	fmt.Println("Seek set to ", newSeekPos)

	_ = json.NewDecoder(r).Decode(&personAgain)
	fmt.Println(personAgain)
}

func decodeSample() {
	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
