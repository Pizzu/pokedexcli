package pokedex

import (
	"errors"

	"github.com/Pizzu/pokedexcli/internal/pokeapi"
)

type PokemonStore interface {
	Add(pokemon pokeapi.PokemonDTO) error
	GetPokemon(name string) (pokeapi.PokemonDTO, error)
	GetAll() ([]pokeapi.PokemonDTO, error)
}

type MapStore struct {
	pokedex map[string]pokeapi.PokemonDTO
}

func NewMapStore() *MapStore {
	return &MapStore{pokedex: make(map[string]pokeapi.PokemonDTO)}
}

func (ms *MapStore) Add(pokemon pokeapi.PokemonDTO) error {
	ms.pokedex[pokemon.Name] = pokemon
	return nil
}

func (ms *MapStore) GetPokemon(name string) (pokeapi.PokemonDTO, error) {
	if pokemon, ok := ms.pokedex[name]; ok {
		return pokemon, nil
	}

	return pokeapi.PokemonDTO{}, errors.New("you have not caught that pokemon")
}

func (ms *MapStore) GetAll() ([]pokeapi.PokemonDTO, error) {
	pokemons := []pokeapi.PokemonDTO{}
	for _, pokemon := range ms.pokedex {
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}
