package main

import (
	"encoding/json"
	"fmt"
	"github.com/foontinz/goBookSolutions/exercises"
	"os"
)

func main() {
	comics, err := exercises.RetrieveAllComics()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	jsonedComics, err := json.MarshalIndent(comics, "", "   ")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	f, err := os.OpenFile("comics.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	_, err = f.Write(jsonedComics)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
