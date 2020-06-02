package main

import (
	"bufio"
	"components"
	"fmt"
	"os"
	"service"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------Welcome To The Game---------------------------")
	fmt.Println("Sample Game Layout for 3 by 3 Game.")
	var firstCol = [3]string{"0", "3", "6"}
	var secondCol = [3]string{"1", "4", "7"}
	var thirdCol = [3]string{"2", "5", "8"}

	fmt.Println(" ___________________________________")
	for i := 0; i < 3; {
		fmt.Println("|           |           |           |")
		fmt.Println("|    ", firstCol[i], "    |    ", secondCol[i], "    |    ", thirdCol[i], "    |")
		fmt.Println("|___________|___________|___________|")
		i = i + 1
	}

	//taking size
	fmt.Print("[Press 2 for 2*2]\n[Press 3 for 3*3]\nEnter size of the board : ")
AGAIN:
	size_of_board, _ := reader.ReadString('\n')
	size_of_board = strings.TrimSpace(size_of_board)
	size, err := strconv.Atoi(size_of_board)
	if err != nil || size > 5 {
		fmt.Print("\nIncorrect Value.\nPlease, try again.\n")
		goto AGAIN
	}

	//taking names and marks
	fmt.Print("Player 1 name : ")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSpace(name1)
	fmt.Print("Player 2 name : ")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)
	fmt.Print("For Player 1 enter 1 to take X and anything else to take O : ")
	mark, _ := reader.ReadString('\n')
	mark = strings.TrimSpace(mark)
	var pl [2]*components.Player
	if mark == "1" {
		pl[0] = components.CreatePlayer(name1, "X")
		pl[1] = components.CreatePlayer(name2, "O")
	} else {
		pl[0] = components.CreatePlayer(name1, "O")
		pl[1] = components.CreatePlayer(name2, "X")
	}

	//starting new game
	newGame := service.NewGameService(pl, size)

	//start playing game
	var res string
	play := true
	fmt.Println(pl[0].Name, " your chance ")
	for {
		fmt.Print(newGame.PrintBoard())
		fmt.Println("Enter position: ")
		pos, _ := reader.ReadString('\n')
		pos = strings.TrimSpace(pos)
		position, err := strconv.Atoi(pos)
		if err != nil {
			fmt.Println("Position should be an integer")
			continue
		}
		if play {
			err, res = newGame.Play(uint8(position))
		} else {
			err, res = newGame.Play(uint8(position))
		}
		if err != nil {
			fmt.Println(err)
			continue
		} else if res == "win" {
			if play {
				fmt.Print(newGame.PrintBoard())
				fmt.Print("Congratulations!! ", pl[0].Name, " has won ")
				break
			} else {
				fmt.Print(newGame.PrintBoard())
				fmt.Print("Congratulations!! ", pl[1].Name, " has won ")
				break
			}
		} else if res == "draw" {
			fmt.Print("Match is Draw!!!!")
			break
		}
		if play {
			fmt.Print("\n", pl[0].Name, " your chance. \n")
		} else {
			fmt.Print("\n", pl[1].Name, " your chance. \n")
		}
		play = !play
	}
}
