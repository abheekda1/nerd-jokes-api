package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Marshal(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func allJokes(w http.ResponseWriter, r *http.Request) {
	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(jokesJSON))
}

func randomJoke(w http.ResponseWriter, r *http.Request) {
	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	type Joke struct {
		Subject    string
		isOneliner bool
		Title      string
		Oneliner   string
		Setup      string
		Punchline  string
	}

	var joke []Joke
	json.Unmarshal([]byte(jokesJSON), &joke)

	json.NewEncoder(w).Encode(joke[rand.Intn(len(joke))])
}

func allJokesBySubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["subject"]

	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	type Joke struct {
		Subject    string
		isOneliner bool
		Title      string
		Oneliner   string
		Setup      string
		Punchline  string
	}

	var joke []Joke
	json.Unmarshal([]byte(jokesJSON), &joke)

	var jokesBySubject []Joke

	for _, subJoke := range joke {
		if subJoke.Subject == key {
			jokesBySubject = append(jokesBySubject, subJoke)
		}
	}
	json.NewEncoder(w).Encode(jokesBySubject)
}

func randomJokeBySubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["subject"]

	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	type Joke struct {
		Subject    string
		isOneliner bool
		Title      string
		Oneliner   string
		Setup      string
		Punchline  string
	}

	var joke []Joke
	json.Unmarshal([]byte(jokesJSON), &joke)

	var jokesBySubject []Joke

	for _, subJoke := range joke {
		if subJoke.Subject == key {
			jokesBySubject = append(jokesBySubject, subJoke)
		}
	}
	json.NewEncoder(w).Encode(jokesBySubject[rand.Intn(len(jokesBySubject))])
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/jokes", allJokes)
	router.HandleFunc("/jokes/random", randomJoke)
	router.HandleFunc("/jokes/{subject}", allJokesBySubject)
	router.HandleFunc("/jokes/random/{subject}", randomJokeBySubject)
	log.Fatal(http.ListenAndServe(":3587", router))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Ready!")
	handleRequests()
}
