package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type user struct {
	Username string `json:"name"`
	Visits   int    `json:"count"`
}

var in InMemory

type InMemory struct {
	usersCOunt map[string]int
	ml         *sync.RWMutex
}

func init() {
	in = InMemory{
		usersCOunt: make(map[string]int),
		ml:         &sync.RWMutex{},
	}
}

func ExposeApiforUserVisits() {

	http.HandleFunc("/user/", userHandler)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error while staring server: %s", err)
	}

}

func userHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		getUserFunc(w, req)
	} else {
		http.Error(w, "method not implmented", http.StatusMethodNotAllowed)
	}

}

func getUserFunc(w http.ResponseWriter, req *http.Request) {

	param := strings.TrimPrefix(req.URL.Path, "/user/")

	in.ml.Lock()
	if cnt, ok := in.usersCOunt[param]; !ok {
		in.usersCOunt[param] = 1
	} else {
		in.usersCOunt[param] = cnt + 1
	}
	in.ml.Unlock()
	retuser := user{
		Username: param,
		Visits:   in.usersCOunt[param],
	}

	json.NewEncoder(w).Encode(retuser)

}
