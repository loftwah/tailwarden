package main

import (
	"log"
	"net/http"
)

// LambdaService is a wrapper around the AWS Lambda API client.
type LambdaService struct{}

// ListFunctions returns a list of all Lambda functions across regions.
func (s *LambdaService) ListFunctions() ([]byte, error) {
	// TODO: Implement function list logic
	return nil, nil
}

// SearchFunctions returns a list of Lambda functions matching the given search criteria.
func (s *LambdaService) SearchFunctions(runtime, tagKey, tagValue, region string) ([]byte, error) {
	// TODO: Implement function search logic
	return nil, nil
}

func main() {
	service := &LambdaService{}

	http.HandleFunc("/functions", func(w http.ResponseWriter, r *http.Request) {
		functions, err := service.ListFunctions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(functions)
	})

	http.HandleFunc("/functions/search", func(w http.ResponseWriter, r *http.Request) {
		runtime := r.URL.Query().Get("runtime")
		tagKey := r.URL.Query().Get("tagKey")
		tagValue := r.URL.Query().Get("tagValue")
		region := r.URL.Query().Get("region")

		functions, err := service.SearchFunctions(runtime, tagKey, tagValue, region)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(functions)
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
