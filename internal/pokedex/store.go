package pokedex

import "github.com/Pizzu/pokedexcli/internal/pokeapi"

type PokemonStore interface {
	Add(pokemon pokeapi.PokemonDTO) error
	Exists(name string) (bool, error)
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

func (ms *MapStore) Exists(name string) (bool, error) {
	_, isPresent := ms.pokedex[name]

	return isPresent, nil
}

func (ms *MapStore) GetAll() ([]pokeapi.PokemonDTO, error) {
	pokemons := []pokeapi.PokemonDTO{}
	for _, pokemon := range ms.pokedex {
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}
