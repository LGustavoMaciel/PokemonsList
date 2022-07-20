package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {

	r := chi.NewRouter()
	r.Get("/pokemons", GetPokemon)

	log.Println("Listening on port 8080...")

	http.ListenAndServe(":8080",r)

	
}


func GetPokemon(w http.ResponseWriter, r *http.Request){
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	
	var responseObject Response
	json.Unmarshal(resData, &responseObject)


	for i := 0; i < len(responseObject.Pokemon); i ++ {
		json.NewEncoder(w).Encode(responseObject.Pokemon[i].Species.Name)
	}
}