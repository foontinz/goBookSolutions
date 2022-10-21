package main

import (
	"bytes"
	"github.com/foontinz/xkcdTool/comicsUtils"
	"log"
	"text/template"
)

const templ = `Num : {{.Num}}
Month : {{.Month}}
Transcript : {{.Transcript}}`

func main() {
	comic := comicsUtils.Comic{
		Num:        0,
		Month:      "July",
		Year:       "2005",
		Transcript: "Something just snapped, something inside of me",
	}

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
