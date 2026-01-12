package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(input []string) int {
	var sum int
	// calculate the inside
	for row := 1; row < len(input)-1; row++ {
		for col := 1; col < len(input[row])-1; col++ {
			var count int
			char := string(input[row][col])
			if char == "@" {
				neighbors := []string{
					string(input[row-1][col]),   // north
					string(input[row+1][col]),   // south
					string(input[row-1][col+1]), // north east
					string(input[row+1][col+1]), // south east
					string(input[row][col+1]),   // east
					string(input[row-1][col-1]), // north west
					string(input[row][col-1]),   // west
					string(input[row+1][col-1]), // south west
				}
				for _, neighbor := range neighbors {
					if neighbor == char {
						count++
					}
				}
				if count < 4 {
					sum++
				}
			}
		}
	}

	// rolls at the corner will always be counted towards the sum total
	if input[0][0] == '@' {
		sum++
	}
	if input[0][len(input)-1] == '@' {
		sum++
	}
	if input[len(input)-1][0] == '@' {
		sum++
	}
	if input[len(input)-1][len(input)-1] == '@' {
		sum++
	}
	// calculate top row without the corner
	for i := 1; i < len(input[0])-1; i++ {
		var count int
		row := 0
		col := i
		char := string(input[0][i])
		if char == "@" {
			neighbors := []string{
				string(input[row+1][col]),   // south
				string(input[row][col-1]),   // west
				string(input[row][col+1]),   // east
				string(input[row+1][col+1]), // south east
				string(input[row+1][col-1]), // south west
			}
			for _, neighbor := range neighbors {
				if neighbor == char {
					count++
				}
			}
			if count < 4 {
				sum++
			}
		}
	}
	// calculate bottom row
	for i := 1; i < len(input[0])-1; i++ {
		var count int
		row := len(input) - 1
		col := i
		char := string(input[row][i])
		if char == "@" {
			neighbors := []string{
				string(input[row][col-1]),   // west
				string(input[row][col+1]),   // east
				string(input[row-1][col+1]), // north east
				string(input[row-1][col-1]), // north west
				string(input[row-1][col]),   // north
			}
			for _, neighbor := range neighbors {
				if neighbor == char {
					count++
				}
			}
			if count < 4 {
				sum++
			}
		}
	}
	// calculate right edge
	for i := 1; i < len(input)-1; i++ {
		var count int
		row := i
		col := len(input[0]) - 1
		char := string(input[i][col])
		if char == "@" {
			neighbors := []string{
				string(input[row][col-1]),   // west
				string(input[row-1][col]),   // north
				string(input[row+1][col]),   // south
				string(input[row-1][col-1]), // north west
				string(input[row+1][col-1]), // south west
			}
			for _, neighbor := range neighbors {
				if neighbor == char {
					count++
				}
			}
			if count < 4 {
				sum++
			}
		}
	}
	// calculate left edge
	for i := 1; i < len(input)-1; i++ {
		var count int
		row := i
		col := 0
		char := string(input[i][col])
		if char == "@" {
			neighbors := []string{
				string(input[row-1][col]),   // north
				string(input[row+1][col]),   // south
				string(input[row-1][col+1]), // north east
				string(input[row+1][col+1]), // south east
				string(input[row][col+1]),   // east
			}
			for _, neighbor := range neighbors {
				if neighbor == char {
					count++
				}
			}
			if count < 4 {
				sum++
			}
		}
	}
	return sum
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		input = append(input, line)
	}
	fmt.Println(part1(input))
}
