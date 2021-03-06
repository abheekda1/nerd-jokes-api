package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Marshal(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func generalJokes(w http.ResponseWriter, r *http.Request) {
	jokesJSON, err := ioutil.ReadFile("jokes.json")
	if err != nil {
		fmt.Println(err)
	}

	type Type1 struct {
		Title    string
		Oneliner string
	}

	type Type2 struct {
		Title     string
		Setup     string
		Punchline string
	}

	type General struct {
		Type1 []Type1
		Type2 []Type2
	}

	type Joke struct {
		General General
	}

	type AllJokes struct {
		Jokes Joke
	}

	var joke AllJokes
	json.Unmarshal([]byte(jokesJSON), &joke)
	jokeType := [2]string{"Type1", "Type2"}
	jokeTypeIndex := jokeType[rand.Intn(2)]

	if jokeTypeIndex == "Type1" {
		index := rand.Intn(len(joke.Jokes.General.Type1))
		randomJoke, _ := Marshal(joke.Jokes.General.Type1[index])
		fmt.Fprintf(w, string(randomJoke))
	}

	if jokeTypeIndex == "Type2" {
		index := rand.Intn(len(joke.Jokes.General.Type2))
		randomJoke, _ := Marshal(joke.Jokes.General.Type2[index])
		fmt.Fprintf(w, string(randomJoke))
	}

	fmt.Println("Endpoint Hit: General Jokes Page")
}

func scienceJokes(w http.ResponseWriter, r *http.Request) {
	jokesJSON, err := ioutil.ReadFile("jokes.json")
	if err != nil {
		fmt.Println(err)
	}

	type Type1 struct {
		Title    string
		Oneliner string
	}

	type Type2 struct {
		Title     string
		Setup     string
		Punchline string
	}

	type Science struct {
		Type1 []Type1
		Type2 []Type2
	}

	type Joke struct {
		Science Science
	}

	type AllJokes struct {
		Jokes Joke
	}

	var joke AllJokes
	json.Unmarshal([]byte(jokesJSON), &joke)
	jokeType := [2]string{"Type1", "Type2"}
	jokeTypeIndex := jokeType[rand.Intn(2)]

	if jokeTypeIndex == "Type1" {
		index := rand.Intn(len(joke.Jokes.Science.Type1))
		randomJoke, _ := Marshal(joke.Jokes.Science.Type1[index])
		fmt.Fprintf(w, string(randomJoke))
	}

	if jokeTypeIndex == "Type2" {
		index := rand.Intn(len(joke.Jokes.Science.Type2))
		randomJoke, _ := Marshal(joke.Jokes.Science.Type2[index])
		fmt.Fprintf(w, string(randomJoke))
	}

	fmt.Println("Endpoint Hit: Science Jokes Page")
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/jokes/random", generalJokes)
	http.HandleFunc("/jokes/science/random", scienceJokes)
	log.Fatal(http.ListenAndServe(":3587", nil))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Ready!")
	handleRequests()
}
