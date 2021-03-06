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

/*func generalJokes(w http.ResponseWriter, r *http.Request) {
	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
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
	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
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
}*/

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

	for _, subJoke := range joke {
		if subJoke.Subject == key {
			fmt.Fprintf(w, "\"Subject\": \"%v\", \"Title\": \"%v\", \"Oneliner\": \"%v\", \"Setup\": \"%v\", \"Punchline\": \"%v\"", subJoke.Subject, subJoke.Title, subJoke.Oneliner, subJoke.Setup, subJoke.Punchline)
		}
	}
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
			jokesBySubject = append(jokesBySubject, fmt.Sprintf("\"Subject\": \"%v\", \"Title\": \"%v\", \"Oneliner\": \"%v\", \"Setup\": \"%v\", \"Punchline\": \"%v\"", subJoke.Subject, subJoke.Title, subJoke.Oneliner, subJoke.Setup, subJoke.Punchline))
		}
	}

	fmt.Fprintf(w, jokesBySubject[rand.Intn(len(jokesBySubject))])

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

	index := rand.Intn(len(joke))
	randomJoke, _ := Marshal(joke[index])
	fmt.Fprintf(w, string(randomJoke))

	fmt.Println("Endpoint Hit: Science Jokes Page")
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
