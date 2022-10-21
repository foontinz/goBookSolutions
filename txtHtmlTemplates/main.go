package main

import (
	"bytes"
	"log"
	"text/template"
)

const templ = `Num : {{.Num}}
Month : {{.Month}}
Year : {{.Year}}
Transcript : {{.Transcript}}`

func main() {
	var comic struct{}
	comicReport, err := template.New("comicReport").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.NewBuffer([]byte{})

	if err = comicReport.Execute(buffer, comic); err != nil {
		log.Fatal(err)
	}
	log.Printf(buffer.String())

}
