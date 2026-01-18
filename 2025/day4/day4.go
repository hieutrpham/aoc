package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func part1(input [][]byte) int {
	var sum int
	input_row_len := len(input)
	input_col_len := len(input[0])
	for row := 0; row < input_row_len; row++ {
		fmt.Println(string(input[row]))
	}
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("\x1b[%dD", input_row_len)
	fmt.Printf("\x1b[%dA", input_col_len)

	saved_input := make([][]byte, input_row_len)
	for row := range saved_input {
		saved_input[row] = make([]byte, input_col_len)
	}
	for row := 0; row < input_row_len; row++ {
		for col := 0; col < input_col_len; col++ {
			var count int
			char := input[row][col]
			if char == '@' {
				// fmt.Println(string(char))
				var neighbors []byte
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						row_x := row + x
						col_y := col + y
						if row_x >= 0 && row_x < input_row_len && col_y >= 0 && col_y < input_col_len {
							if x == y && x == 0 && y == 0 {
								continue
							}
							neighbors = append(neighbors, input[row_x][col_y])
						}
					}
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
	for row := 0; row < len(saved_input); row++ {
		for col := 0; col < len(saved_input[0]); col++ {
			if saved_input[row][col] == 'x' {
				input[row][col] = saved_input[row][col]
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

	//part2
	sum := 0
	for {
		count := part1(grid2d)
		if count <= 0 {
			break
		}
		sum += count
	}
	fmt.Print("\x1b[H\x1b[2J")
	// fmt.Println(sum)
}
