package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("not enough arguments")
	}

	start := time.Now()

	//var wg sync.WaitGroup

	ch := make(chan string,len(os.Args[1:])) // memory is already allocated in buffered channel

	for _, url := range os.Args[1:] {

		//wg.Add(1)
		go func(url string) {
			log.Printf("url is %s", url)
			resp, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprintf("error for url %s : %v", url,err)
				return
			}
			ch <- fmt.Sprintf("statuscode for url %s is: %d", url, resp.StatusCode)

			//wg.Done()
		}(url)

	}
	//wg.Wait()

	for i := cap(ch);i>0;i-- {
		fmt.Println(<- ch)
	}

	log.Printf("time Duration is: %s", time.Since(start))
}

