package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status int
}

func getStatusOfUrls() {

	urls := []string{
		"https://google.com",
		"https://example.com",
		"https://reddit.com",
		"https://not-a-valid-url.com",
	}

	statusChan := make(chan urlStatus)
	for _, url := range urls {

		go getStatus(url, statusChan)
	}

	for us := range statusChan {
		fmt.Println(us.url, ":", us.status)
	}

	close(statusChan)

}

func getStatus(url string, statuschan chan urlStatus) {

	urlStruct := urlStatus{
		url: url,
	}
	var status int
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while calling the url:", url, err)
	}

	if resp != nil {
		status = resp.StatusCode
		defer resp.Body.Close()
	} else {
		status = http.StatusNotFound
	}

	urlStruct.status = status
	statuschan <- urlStruct
}
