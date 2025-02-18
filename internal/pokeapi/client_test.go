package pokeapi

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Pizzu/pokedexcli/internal/pokecache"
)

type MockRoundTripper struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func NewMockRoundTripper(fn func(*http.Request) (*http.Response, error)) *MockRoundTripper {
	return &MockRoundTripper{
		roundTripFunc: fn,
	}
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTripFunc(req)
}

func TestListPokemonsWithinLocation(t *testing.T) {
	t.Run("successful pokemon list", func(t *testing.T) {
		mockTransport := NewMockRoundTripper(
			func(req *http.Request) (*http.Response, error) {
				// Verify correct URL being called
				if !strings.Contains(req.URL.String(), "location-area/test-area") {
					t.Errorf("unexpected URL: %s", req.URL.String())
				}

				json := `{
                    "id": 1,
                    "location": {
                        "name": "test-location",
                        "url": "https://example.com/location/1"
                    },
                    "name": "test-area",
                    "pokemon_encounters": [
                        {
                            "pokemon": {
                                "name": "pikachu",
                                "url": "https://example.com/pokemon/25"
                            }
                        }
                    ]
                }`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader([]byte(json))),
				}, nil
			},
		)

		client := &Client{
			httpClient: http.Client{Transport: mockTransport},
			cache:      pokecache.NewCache(5 * time.Second),
		}

		result, err := client.ListPokemonsWithinLocation("test-area")

		if err != nil {
			t.Errorf("Error %s", err.Error())
			return
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
		if result.Name != "test-area" {
			t.Errorf("Expected name 'test-area', got %s", result.Name)
		}
		if result.Location.Name != "test-location" {
			t.Errorf("Expected location name 'test-location', got %s", result.Location.Name)
		}
		if len(result.PokemonEncounters) != 1 {
			t.Errorf("Expected 1 pokemon encounter, got %d", len(result.PokemonEncounters))
		}
		if result.PokemonEncounters[0].Pokemon.Name != "pikachu" {
			t.Errorf("Expected pokemon name 'pikachu', got %s", result.PokemonEncounters[0].Pokemon.Name)
		}
	})
}
