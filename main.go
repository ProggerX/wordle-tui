package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

//go:embed words.txt
var str string

func main() {
	words := strings.Split(str, "\n")
	word := words[rand.Intn(5757)]
	color.White("You have 5 tries to guess:")
	for range 5 {
		var guess string
		fmt.Scan(&guess)
		if len(guess) != 5 {
			color.White("You need to guess 5-letter words!")
			os.Exit(0)
		}
		if guess == word {
			color.White("You guessed correctly!")
			os.Exit(0)
		}
		var count = make(map[string]int, 0)
		for i, v := range guess {
			if string(v) == string(word[i]) {
				fmt.Print(color.GreenString(string(v)))
				count[string(v)]++
			} else if strings.Count(word, string(v))-count[string(v)] > 0 {
				fmt.Print(color.YellowString(string(v)))
				count[string(v)]++
			} else {
				fmt.Print(color.BlackString(string(v)))
			}
		}
		fmt.Println()
	}
	color.White("Sorry, you didn't guess right!")
	color.White("Word was: " + word)
}
