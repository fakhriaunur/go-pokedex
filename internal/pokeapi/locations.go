package pokeapi

import "encoding/json"

type Location struct {
	Name string `json:"name"`
}

type ApiResponse struct {
	Results  []Location `json:"results"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
}

func UnmarshalLocations(data []byte) (*ApiResponse, error) {
	var apiResponse ApiResponse
	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
