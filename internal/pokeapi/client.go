package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetAndParse(url string) (Response, error) {

	var results Response

	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&results); err != nil {
		return results, err
	}
	return results, nil
}
