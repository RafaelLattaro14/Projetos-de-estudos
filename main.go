package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

// main is the entry point of the guessing game.
// It generates a random number between 0 and 100, and prompts the user to guess it.
// The game ends when the user guesses the correct number or after 10 attempts.
func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println(
		"Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")

	// Generate a random number between 0 and 100
	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	// Loop until the user guesses the correct number or runs out of attempts
	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		// Parse the user's guess into an integer
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número de 0 a 100!")
			return
		}

		// Compare the user's guess with the generated number
		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			// User guessed correctly
			fmt.Printf("Parabéns! O número era: %d\n"+
				"Você acertou em %d tentativas\n"+
				"Essas foram suas tentativas: %v\n", x, i+1, chutes[:i])
			return
		}

		// Store the user's guess for future reference
		chutes[i] = chuteInt
	}

	// User ran out of attempts
	fmt.Printf("Infelizmente você não acertou o número, que era: %d\n"+
		"Você teve 10 tentativas. \n"+
		"Essas foram suas tentativas: %v\n", x, chutes)
}
