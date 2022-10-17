package exercises

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Comics struct {
	Items []Comic `json:"items"`
}
type Comic struct {
	Num        int    `json:"num"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Transcript string `json:"transcript"`
}

const ComicUrl = "https://xkcd.com/%v/info.0.json"

func RetrieveAllComics() (Comics, error) {
	ComicsArr := []Comic{}
	for page := 1; page < 2046; page++ {
		resp, err := http.Get(fmt.Sprintf(ComicUrl, page))
		if err != nil {
			return Comics{}, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return Comics{}, fmt.Errorf("request to xkcd failed with %s, %v", resp.Status, page)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			return Comics{}, err
		}

		var comic Comic
		if err := json.Unmarshal(body, &comic); err != nil {
			resp.Body.Close()
			return Comics{}, err
		}
		ComicsArr = append(ComicsArr, comic)

	}

	return Comics{ComicsArr}, nil
}
