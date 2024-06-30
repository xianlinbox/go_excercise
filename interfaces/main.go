package main

import (
	"fmt"
	"os"
)

type bot interface {
	getGreeting() string;
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	fmt.Println(os.Args)

	fileReader := fileReader{}
	fileReader.readFromFile(os.Args[1])
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}