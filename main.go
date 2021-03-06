package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
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

	var jokesBySubject []string

	for _, subJoke := range joke {
		jokesBySubject = append(jokesBySubject, fmt.Sprintf("{\n \"Subject\": \"%v\", \n \"Title\": \"%v\", \n \"Oneliner\": \"%v\", \n \"Setup\": \"%v\", \n \"Punchline\": \"%v\"\n}", subJoke.Subject, subJoke.Title, subJoke.Oneliner, subJoke.Setup, subJoke.Punchline))
	}

	fmt.Fprintf(w, jokesBySubject[rand.Intn(len(jokesBySubject))])

	/*index := rand.Intn(len(joke))
	randomJoke, _ := Marshal(joke[index])
	fmt.Fprintf(w, string(randomJoke))*/
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

	var jokesBySubject string

	for _, subJoke := range joke {
		if subJoke.Subject == key {
			jokesBySubject += fmt.Sprintf("\n {\n  \"Subject\": \"%v\", \n  \"Title\": \"%v\", \n  \"Oneliner\": \"%v\", \n  \"Setup\": \"%v\", \n  \"Punchline\": \"%v\"\n },", subJoke.Subject, subJoke.Title, subJoke.Oneliner, subJoke.Setup, subJoke.Punchline)
		}
	}
	fmt.Fprintf(w, "["+strings.TrimSuffix(jokesBySubject, ",")+"\n]")
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

	var jokesBySubject []string

	for _, subJoke := range joke {
		if subJoke.Subject == key {
			jokesBySubject = append(jokesBySubject, fmt.Sprintf("{\n \"Subject\": \"%v\", \n \"Title\": \"%v\", \n \"Oneliner\": \"%v\", \n \"Setup\": \"%v\", \n \"Punchline\": \"%v\"\n}", subJoke.Subject, subJoke.Title, subJoke.Oneliner, subJoke.Setup, subJoke.Punchline))
		}
	}

	fmt.Fprintf(w, jokesBySubject[rand.Intn(len(jokesBySubject))])

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
