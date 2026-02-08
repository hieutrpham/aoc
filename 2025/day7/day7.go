package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print_array[T any](input []T) {
	for i := range input {
		fmt.Println(input[i])
	}
}

func part1(input [][]string) int {
	row := len(input)
	col := len(input[0])
	count := 0
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			curr_char := input[y][x]
			if curr_char == "S" {
				input[y+1][x] = "|"
			}
			if curr_char == "^" && input[y-1][x] == "|" {
				input[y+1][x-1] = "|"
				input[y+1][x+1] = "|"
				count++
			}
			if curr_char == "." && y > 0 && input[y-1][x] == "|" {
				input[y][x] = "|"
			}
		}
	}
	return count
}

func get_input() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	var input [][]string
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}
	return input
}

type Coord struct {
	x int
	y int
}

type Node struct {
	coord Coord
	left  *Node
	right *Node
}

func main() {
	input := get_input()
	part1(input)
}
