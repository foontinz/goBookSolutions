package comicsUtils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Comic struct {
	Num        int    `json:"num"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Transcript string `json:"transcript"`
}

const ComicUrl = "https://xkcd.com/%v/info.0.json"

func DownloadAllComics() ([]Comic, error) {
	var ComicsArr []Comic
	for page := 1; page < 2046; page++ {
		resp, err := http.Get(fmt.Sprintf(ComicUrl, page))
		if err != nil {
			return ComicsArr, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			log.Printf("request to xkcd failed with %s, %v", resp.Status, page)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			return ComicsArr, err
		}

		var comic Comic
		if err := json.Unmarshal(body, &comic); err != nil {
			resp.Body.Close()
			return ComicsArr, err
		}
		ComicsArr = append(ComicsArr, comic)

	}

	return ComicsArr, nil
}

func SaveComics(file string) (bool, error) {
	comics, err := DownloadAllComics()
	if err != nil {
		return false, err
	}
	jsonedComics, err := json.MarshalIndent(comics, "", "    ")
	if err != nil {
		return false, err
	}

	f, err := os.Open(file)
	if err != nil {
		return false, err
	}
	_, err = f.Write(jsonedComics)
	if err != nil {
		return false, err
	}
	return true, nil
}

func LoadComics(file string) ([]Comic, error) {
	data, err := os.ReadFile(file)
	var comics []Comic

	if err != nil {
		return comics, err
	}
	err = json.Unmarshal(data, &comics)
	if err != nil {
		return comics, err
	}

	return comics, nil
}
