package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Joke struct {
	ID   string `json:"id"`
	Text string `json:"value"`
}

func sayhi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "im saying hello!!!")
}

func gjokes(w http.ResponseWriter, r *http.Request) {
	var jokes []Joke
	ids := make(map[string]bool)

	for len(jokes) < 25 {
		joke, err := getJoke()
		if err != nil {
			fmt.Fprint(w, "hubo un error")
		}
		if !ids[joke.ID] {
			ids[joke.ID] = true
			jokes = append(jokes, joke)
		}
	}
	//fmt.Println(jokes)
	jsonBytes, err := json.Marshal(jokes)
	if err != nil {
		fmt.Println("Error al convertir a Json", err)
		return
	}
	fmt.Fprint(w, string(jsonBytes))
}

func getJoke() (Joke, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return Joke{}, err
	}
	return joke, nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sayhi)
	router.HandleFunc("/getlist", gjokes)
	handle := cors.Default().Handler(router)
	fmt.Println("Server is en port 3000")
	log.Fatal(http.ListenAndServe(":3000", handle))
}
