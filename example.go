package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// API response structure
type Response struct {
	Status  bool        `json:"status"`
	Creator string      `json:"creator"`
	Result  interface{} `json:"result"`
}

func main() {
	apiKey := "YOUR_API_KEY" // 🔑 Replace or load from env
	endpoint := "https://discardapi.dpdns.org/api/endpoint"

	url := fmt.Sprintf("%s?apikey=%s", endpoint, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Invalid JSON: %v", err)
	}

	fmt.Println("✅ API Response:")
	fmt.Printf("Status: %v\nCreator: %s\nResult: %v\n", data.Status, data.Creator, data.Result)
}
