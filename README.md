# Nerd Jokes API
A nerdy API for categorized jokes, heavily focused on science puns!

## About
I was looking around for jokes APIs, but couldn't find any that had science jokes that I could request for specifically. As I had just started learning golang, I decided to waste an entire day making a crappy, useless API that has ugly code and probably has tons of extra garbage - but it works!

## Usage
To use it, just clone the repository and run either `go run main.go` for a quick test or `go build main.go` and `./main` for use in production. By default, it runs on port 3587 (as I use a reverse proxy for that port to be accesed by a certain domain) but this can be modified in the `main.go` file.

### Requesting jokes
You can take a look at the endpoints in the `main.go` file, but here are the main ones: `/jokes/random`, `/jokes/random/{subject}`, `/jokes/{subject}` and `/jokes`. These only work with GET requests.

## Adding jokes
You can send a POST request to the `/addJoke` endpoint with only the fields that have content (take a look at `jokes.json` for the format). For example, you can have
```
{
  apikey: 'specified on start',
  subject: 'subject',
  setup: 'funny joke',
  punchline: 'lol lol'
}
 ```
 instead of specifying each field if it isn't used. Postman is easy to use for this but cURL also works.
 > Note: the API Key is specified on the launching of the program. For example `./main [API Key here]`.
