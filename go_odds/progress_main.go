package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// JobProgress is request structure for Migration Stats API
type JobProgress struct {
	TotalFiles    int64 `json:"total_files"`
	TotalSize     int64 `json:"total_size"`
	FilesMigrated int64 `json:"files_migrated"`
	SizeMigrated  int64 `json:"size_migrated"`
	FilesFailed   int64 `json:"files_failed"`
	Speed         int   `json:"speed"`
	Eta           int   `json:"eta"`
}

// JobProgressResponse is response structure for Migration Stats API
type JobProgressResponse struct {
	Status bool `json:"status"`
}

func updateJobProgressHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the job ID from the URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	jobID := parts[len(parts)-2]

	fmt.Println("job id recieved :", jobID)
	// Parse the request body into a JobProgress struct
	var request JobProgress
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// TODO: Implement logic to update job progress based on the jobID and request data

	// Simulating a successful response
	response := JobProgressResponse{
		Status: true,
	}

	// Set the response content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}

func proc_main() {
	http.HandleFunc("/api/v1/jobs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/status_update") {
			updateJobProgressHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Println("Server started on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
