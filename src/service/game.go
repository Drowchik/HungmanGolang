package service

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"unicode/utf8"
)

const MaxErrors = 6

func Start() {
	words, err := ReadWordFromFile("src/data/book.txt")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("=== Новая игра ===")
		word := words[rand.IntN(len(words))]
		hidden := strings.Repeat("*", utf8.RuneCountInString(word))
		errors := 0
		used := make(map[rune]bool)
		for {
			if errors >= MaxErrors {
				fmt.Printf("Ты проиграл! Слово было: %s\n", word)
				break
			}

			fmt.Println(RenderState(hidden, errors))
			char := string(ReadValidatedInput(used))

			if strings.Contains(word, char) {
				hidden = revealLetters(word, hidden, char)
				if hidden == word {
					fmt.Printf("Поздравляем! Слово угадано: %s\n", word)
					break
				}
			} else {
				errors++
			}
		}
		fmt.Println("Хотите ли вы продолжить? Для завершения нажмити n/N")
		var input string
		fmt.Scanln(&input)
		if input == "n" || input == "N" {
			fmt.Println("Приходите еще!")
			break
		}
	}

}

func revealLetters(word, current string, guess string) string {
	result := []rune(current)
	i := 0
	for _, ch := range word {
		if string(ch) == guess {
			result[i] = ch
		}
		i++
	}
	return string(result)
}
