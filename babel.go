package main

import (
	"fmt"
	"math/rand"
	"time"
)

func WordGenerator(length int) string { //function designed to generate a random word after intaking the length of the word.
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	word := make([]byte, length)
	for i := range word {
		word[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(word)
}

func main() {

	var a = rand.Intn(300)
	var b [301]string

	for i := 0; i < a; i++ {
		l := rand.Intn(10) + 2
		b[i] = WordGenerator(l)
		fmt.Print(b[i], " ")
	}

	fmt.Println("This will be a book of babel simulation")
	fmt.Println("", a)

}
