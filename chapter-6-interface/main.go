package main

import "fmt"

type bot interface {
	getGreetings () string
}

type englishBot struct{}
type spanishBot struct{}

func main () {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreetings () string {
	return "Hello"
}

func (sb spanishBot) getGreetings () string {
	return "Hola!"
}

func printGreeting (b bot) {
	fmt.Println(b.getGreetings())
}