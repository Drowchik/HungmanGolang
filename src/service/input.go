package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadValidatedInput(used map[rune]bool) rune {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите одну русскую букву: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len([]rune(input)) != 1 {
			fmt.Println("Ошибка: нужно ввести одну букву.")
			continue
		}

		r := []rune(input)[0]

		if !isRussianLetter(r) {
			fmt.Println("Ошибка: только русские буквы.")
			continue
		}

		if used[r] {
			fmt.Println("Вы уже вводили эту букву.")
			continue
		}

		used[r] = true
		return r
	}
}

func isRussianLetter(r rune) bool {
	return (r >= 'а' && r <= 'я') || r == 'ё'
}
