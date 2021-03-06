package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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

func addJoke(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	jokesJSON, err := ioutil.ReadFile("static/jokes.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	type Joke struct {
		APIKey    string
		Subject   string
		Title     string
		Oneliner  string
		Setup     string
		Punchline string
	}

	var joke Joke
	json.Unmarshal(reqBody, &joke)

	if joke.APIKey != os.Args[1] {
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	jokesJSONString := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(string(jokesJSON), "\n"), "]"), "\n") + ",\n  " + fmt.Sprintf("{\n    \"subject\": \"%v\", \n    \"title\": \"%v\", \n    \"oneliner\": \"%v\", \n    \"setup\": \"%v\", \n    \"punchline\": \"%v\"\n  }\n]", joke.Subject, joke.Title, joke.Oneliner, joke.Setup, joke.Punchline)

	os.Remove("static/jokes.json")

	f, err := os.OpenFile("static/jokes.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(jokesJSONString); err != nil {
		panic(err)
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/jokes", allJokes).Methods("GET")
	router.HandleFunc("/jokes/random", randomJoke).Methods("GET")
	router.HandleFunc("/jokes/{subject}", allJokesBySubject).Methods("GET")
	router.HandleFunc("/jokes/random/{subject}", randomJokeBySubject).Methods("GET")
	router.HandleFunc("/addJoke", addJoke).Methods("POST")
	log.Fatal(http.ListenAndServe(":3587", router))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Ready!")
	handleRequests()
}
