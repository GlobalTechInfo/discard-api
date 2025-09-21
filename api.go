package discardapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Creator string      `json:"creator"`
	Result  interface{} `json:"result"`
}

func Get(endpoint, apikey string) (*Response, error) {
	url := fmt.Sprintf("%s?apikey=%s", endpoint, apikey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
