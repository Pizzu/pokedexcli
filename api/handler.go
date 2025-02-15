package api

import (
	"encoding/json"
	"net/http"
)

func GetAllLocationArea(url string) (LocationDTO, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationDTO{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return LocationDTO{}, err
	}

	defer res.Body.Close()

	var locationDTO LocationDTO
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationDTO)

	if err != nil {
		return locationDTO, err
	}

	return locationDTO, nil
}
