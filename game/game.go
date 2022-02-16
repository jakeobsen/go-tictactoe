// Copyright 2022 Morten Jakobsen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package game

import (
	"fmt"
	"strconv"
)

var (
	pos            = []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	current_player = "x"
	next_move      = ""
	winning_player = ""
	total_plays    = 0
)

// Main exported game function

func Play() {
	printGameField()

	for {
		askPlayerForMove()
		tryPerformPlayerMove()
		printGameField()
		checkForWinner()

		if winning_player != "" {
			if winning_player == "draw" {
				fmt.Println("It's a draw!")
			} else {
				fmt.Println("Winner is " + winning_player + "!")
			}
			break
		}
	}

}

// Game logic below

func checkForWinner() {
	// This function checks all possible combinations to win tic tac toe
	// It also checks if the game is a draw

	if pos[1] == pos[2] && pos[2] == pos[3] {
		winning_player = pos[1]
	} else if pos[4] == pos[5] && pos[5] == pos[6] {
		winning_player = pos[4]
	} else if pos[7] == pos[8] && pos[8] == pos[9] {
		winning_player = pos[7]
	} else if pos[7] == pos[4] && pos[4] == pos[1] {
		winning_player = pos[7]
	} else if pos[8] == pos[5] && pos[5] == pos[2] {
		winning_player = pos[8]
	} else if pos[9] == pos[6] && pos[6] == pos[3] {
		winning_player = pos[9]
	} else if pos[7] == pos[5] && pos[5] == pos[3] {
		winning_player = pos[7]
	} else if pos[9] == pos[5] && pos[5] == pos[1] {
		winning_player = pos[9]
	} else if total_plays == 9 {
		winning_player = "draw"
	} else {
		winning_player = ""
	}
}

func tryPerformPlayerMove() {
	// This function tries to perform the players desired move
	// If the move is possible, then the move is carried out

	move, err := strconv.ParseInt(next_move, 0, 8)
	if err == nil && move > 0 && move < 10 && pos[move] == next_move {
		pos[move] = current_player
		total_plays++
		switchPlayer()
	} else {
		fmt.Println("Bad move, try again.")
	}
}

func switchPlayer() {
	// Toggle between player x and o

	if current_player == "x" {
		current_player = "o"
	} else {
		current_player = "x"
	}
}

func askPlayerForMove() {
	// Get players input for their next move
	// move validation is carried out in the tryPerformPlayerMove function

	fmt.Print("Player " + current_player + ", make a move: ")
	_, err := fmt.Scanln(&next_move)
	if err != nil {
		fmt.Println("Error!")
		return
	}
}

func printGameField() {
	fmt.Println("┌───┬───┬───┐")
	fmt.Println("│ " + pos[7] + " │ " + pos[8] + " │ " + pos[9] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + pos[4] + " │ " + pos[5] + " │ " + pos[6] + " │")
	fmt.Println("├───┼───┼───┤")
	fmt.Println("│ " + pos[1] + " │ " + pos[2] + " │ " + pos[3] + " │")
	fmt.Println("└───┴───┴───┘")
}
