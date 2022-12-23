package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/eightball", magic8ball)
	http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	info := request.URL.Path[1:]
	fmt.Fprintf(writer, "<h1>hello %s</h1>", info)
}

func magic8ball(writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().UnixNano())
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Fprintf(writer, "<h1>magic 8 ball 🎱 says: %s</h1>", answers[rand.Intn(len(answers))])
}
