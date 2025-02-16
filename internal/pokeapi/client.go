package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Pizzu/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(5 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

func (c *Client) ListLocations(pageURL *string) (LocationDTO, error) {
	var locationDTO LocationDTO
	if rawData, ok := c.cache.Get(*pageURL); ok {
		err := json.Unmarshal(rawData, &locationDTO)

		if err != nil {
			return locationDTO, err
		}

		return locationDTO, nil
	} else {
		req, err := http.NewRequest("GET", *pageURL, nil)

		if err != nil {
			return LocationDTO{}, err
		}

		req.Header.Set("Content-Type", "application/json")

		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationDTO{}, err
		}

		if res == nil || res.StatusCode != http.StatusOK {
			return LocationDTO{}, fmt.Errorf("non-OK HTTP status: %s", res.Status)
		}

		defer res.Body.Close()

		// Save url and res to cache
		rawData, err := io.ReadAll(res.Body)

		
		if err != nil {
			return LocationDTO{}, err
		}

		
		err = json.Unmarshal(rawData, &locationDTO)
		
		if err != nil {
			return locationDTO, err
		}

		c.cache.Add(*pageURL, rawData)

		return locationDTO, nil
	}
}
