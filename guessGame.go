package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func input() (int, error) {
	line := bufio.NewReader(os.Stdin)
	inp, e := line.ReadString('\n')
	inp = strings.TrimSpace(inp)
	num, e := strconv.Atoi(inp)
	if e != nil {
		return 0, errors.New("You have to input an integer")
	}
	if num < 1 || num > 50 {
		return 0, errors.New("Number should be from 1 to 50")
	}
	return num, nil
}

func checknumber(guess int, randomNumber int) bool {
	if guess < randomNumber {
		fmt.Println("Low Guess")
		return false
	} else if guess > randomNumber {
		fmt.Println("High Guess")
		return false
	} else {
		fmt.Println("Perfect Guess")
		return true
	}
}

func main() {
	var gameStatus bool
	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(50) + 1
	fmt.Println("Enter your guess from 1 to 50")
	for attemptsLeft := 5; attemptsLeft > 0; {
		fmt.Println("Attempts left ", attemptsLeft)
		fmt.Print("Enter your guess : ")
		guess, e := input()
		if e != nil {
			fmt.Println("Error: ", e)
			continue
		}
		attemptsLeft--
		gameStatus = checknumber(guess, randomNumber)
		if gameStatus {
			fmt.Println("You took", 5-attemptsLeft, "chances to guess")
			break
		}
		if !gameStatus && attemptsLeft == 0 {
			fmt.Println("The correct answer is:", randomNumber)
		}
	}
}
