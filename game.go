package mastermindgo

import (
	"fmt"
	"math/rand"
	"strings"
)

const codeLength int = 4
const maxGuesses int = 10

func Play() {
	codeChars := []string{"R", "G", "B", "Y"}
	code := generateCode(codeChars)

	for guessesCount := 0; ; {
		if guessesCount >= maxGuesses {
			fmt.Printf("Unlucky! You ran out of guesses. The code was %s\n", code)
			break
		}

		var guess [codeLength]string

		fmt.Printf("You have %d/%d guesses remaining.\n", maxGuesses-guessesCount, maxGuesses)
		fmt.Printf("Enter your guess (space separated, valid characters %v): ", codeChars)
		fmt.Scanln(&guess[0], &guess[1], &guess[2], &guess[3])

		correct_position, incorrect_position, err := evaluateGuess(guess, code)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			guessesCount++
		}

		if correct_position == codeLength {
			fmt.Printf("Congrats! The code was %s\n", code)
			break
		}

		fmt.Printf("Correct position: %d | Incorrect position: %d\n", correct_position, incorrect_position)
	}

}

// Generate a random code of length codeLength from the defined codeChars
func generateCode(codeChars []string) [codeLength]string {
	code := [codeLength]string{}
	for i := 0; i < codeLength; i++ {
		random_char_index := rand.Intn(len(codeChars))
		code[i] = codeChars[random_char_index]
	}

	return code
}

func evaluateGuess(guess, code [codeLength]string) (int, int, error) {
	var correct_position int
	var incorrect_position int

	for guessIdx, guessChar := range guess {
		if len(guessChar) != 1 {
			return 0, 0, fmt.Errorf("guess must consist of 4 space-separated characters")
		}

		if strings.EqualFold(guessChar, code[guessIdx]) {
			correct_position++
			continue
		}

		for _, codeChar := range code {
			if strings.EqualFold(guessChar, codeChar) {
				incorrect_position++
				break
			}
		}
	}

	return correct_position, incorrect_position, nil
}
