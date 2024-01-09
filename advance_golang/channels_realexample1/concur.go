package concurrency

import "time"

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

type WebsiteChecker func(string) bool

func mockWebsitechecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

type result struct {
	string
	bool
}

// WebsiteChecker is the "Dependency injection here"
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string){
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i:=0; i<len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}