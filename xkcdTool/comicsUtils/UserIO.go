package comicsUtils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func validateInputChapter(input string, dataArr []Comic) bool {
	num, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	if !(num < len(dataArr)) || !(num >= 0) {
		return false
	}

	return true
}
func SearchTranscriptInterface(file string) bool {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	data, err := LoadComics(file)

	if err != nil {
		return false
	}

	if string(line) == "exit" {
		return false
	}

	if err != nil || !validateInputChapter(string(line), data) {
		SearchTranscriptInterface(file)
		return true
	}

	chapter, _ := strconv.Atoi(string(line))

	fmt.Printf("%v\n", data[chapter].Transcript)
	return true

}
