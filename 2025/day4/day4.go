package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(input [][]byte) int {
	var sum int
	// calculate the inside
	saved_input := make([][]byte, len(input))
	for row := range saved_input {
		saved_input[row] = make([]byte, len(input[0]))
	}
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
					saved_input[row][col] = 'x'
				}
			}
		}
	}

	// rolls at the corner will always be counted towards the sum total
	if input[0][0] == '@' {
		sum++
		saved_input[0][0] = 'x'
	}
	if input[0][len(input)-1] == '@' {
		saved_input[0][len(input)-1] = 'x'
		sum++
	}
	if input[len(input)-1][0] == '@' {
		sum++
		saved_input[len(input)-1][0] = 'x'
	}
	if input[len(input)-1][len(input)-1] == '@' {
		sum++
		saved_input[len(input)-1][len(input)-1] = 'x'
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
				saved_input[0][i] = 'x'
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
				saved_input[row][i] = 'x'
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
				saved_input[i][col] = 'x'
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
				saved_input[i][col] = 'x'
				sum++
			}
		}
	}
	for row := 0; row < len(saved_input); row++ {
		for col := 0; col < len(saved_input[0]); col++ {
			if saved_input[row][col] == 'x' {
				input[row][col] = saved_input[row][col]
			}
		}
	}
	// fmt.Println(sum)
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
	row_count := len(input)
	col_count := len(input[0])
	grid2d := make([][]byte, row_count)
	for row := range grid2d {
		grid2d[row] = make([]byte, col_count)
	}
	for row := 0; row < row_count; row++ {
		for col := 0; col < row_count; col++ {
			grid2d[row][col] = input[row][col]
		}
	}

	// part1
	// fmt.Println("part 1:", part1(grid2d))

	// part2
	sum := 0
	for {
		count := part1(grid2d)
		if count <= 0 {
			break
		}
		sum += count
	}
	fmt.Println(sum)
}
