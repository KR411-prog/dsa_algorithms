package main

// https://dev.to/digi0ps/making-concurrent-api-requests-in-go-4jig
import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

func getComic(comicID int) (comic Comic, err error) {
	url := fmt.Sprintf(baseXkcdURL, comicID)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return Comic{}, err
	}
	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return Comic{}, err
	}
	return comic, err
}

func main() {

	start := time.Now()

	defer func() {
		fmt.Println("Execution time: ", time.Since(start))
	}()
	comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//comicMap := make(map[int]*Comic, len(comicsNeeded))

	comicChan := make(chan Comic)

	wg := sync.WaitGroup{}

	for _, id := range comicsNeeded {

		wg.Add(1)
		go func(id int) {
			comic, err := getComic(id)

			if err != nil {
				return
			}
			//comicMap[id] = comic
			comicChan <- comic
			wg.Done()
		}(id)

	}
	var comics []Comic

	for i:=0; i<len(comicsNeeded); i++ {
		comics = append(comics, <- comicChan)
		fmt.Printf("Fetched comic %d with title %v\n", comicsNeeded[i], comics[i].Title)

	}
	wg.Wait()
}
