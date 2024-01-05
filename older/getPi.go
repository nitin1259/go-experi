package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// https://gita-api.vercel.app/docs#/default

type apiUrlRespStatus struct{
	url string
	resp map[string]interface{}
}

func GetPi_main() {
	

	urls := []string{
		"https://gita-api.vercel.app/odi/verse/1",
		"https://gita-api.vercel.app/tel/verse/1/1",
		"https://gita-api.vercel.app/",
	}

	ch := make(chan apiUrlRespStatus )

	for _,apiUrl := range urls{

		go doAPICall(apiUrl, ch)
		for resp := range ch{
			fmt.Printf("url: %s, resp: %v", resp.url, resp.resp)
		}
	}

	close(ch)
}

func doAPICall(apiUrl string, ch chan apiUrlRespStatus){

	respstatus := apiUrlRespStatus{
		url: apiUrl,
	} 

	response, err := http.Get(apiUrl)

		if err!=nil{

			log.Fatal("Error while making http call, url:", apiUrl)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)

		if err!=nil{
			fmt.Println("Error while pasing body for url:", apiUrl, err)
		}

		var apiResp map[string]interface{}
		// fmt.Println(string(body))

		err = json.Unmarshal(body, &apiResp)
		if err!=nil{
			fmt.Println("Error while parsing error:", err)
		}

		respstatus.resp = apiResp

		ch <- respstatus
}