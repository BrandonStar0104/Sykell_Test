package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"sykell/test/models"
	"sykell/test/utils"
)

// Global in-memory store for results
var (
	results   = make(map[string]models.ProcessResult)
	resultsMu sync.Mutex
)

func ProcessURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request models.ProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(request.URL, "http://") && !strings.HasPrefix(request.URL, "https://") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	result := utils.AnalyzeURL(request.URL)
	resultsMu.Lock()
	results[request.URL] = result
	resultsMu.Unlock()

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	resultsMu.Lock()
	defer resultsMu.Unlock()

	var urls []models.ProcessResult
	for _, result := range results {
		urls = append(urls, result)
	}

	response, err := json.Marshal(urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	url := strings.TrimPrefix(r.URL.Path, "/result/")
	if url == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	resultsMu.Lock()
	result, exists := results[url]
	resultsMu.Unlock()

	if !exists {
		http.Error(w, "Result not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func StopProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request models.ProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultsMu.Lock()
	delete(results, request.URL)
	resultsMu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Processing stopped for URL: " + request.URL))
}
