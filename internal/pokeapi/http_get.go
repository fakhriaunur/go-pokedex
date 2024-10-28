package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

func FetchData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch data")
	}

	return io.ReadAll(res.Body)
}
