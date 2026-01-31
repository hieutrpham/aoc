package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_num(char string) int {
	num, ok := strconv.Atoi(char)
	if ok != nil {
		fmt.Println("num:", num)
		panic("ERR atoi")
	}
	return num
}

func part1(clean_input []string, stride int, row int) int {
	total := 0
	var answer int
	for x := 0; x < stride; x++ {
		op := clean_input[(row-1)*stride+x]
		if op == "*" {
			answer = 1
		} else {
			answer = 0
		}
		for y := 0; y < row-1; y++ {
			value := clean_input[y*stride+x]
			if op == "*" {
				answer *= get_num(value)
			} else {
				answer += get_num(value)
			}
		}
		total += answer
	}
	return total
}

func main() {
	var input [][]string
	row := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		input = append(input, strings.Split(line, " "))
		row++
	}

	var clean_input []string
	for row := 0; row < len(input); row++ {
		for item := 0; item < len(input[row]); item++ {
			value := input[row][item]
			if len(value) != 0 {
				clean_input = append(clean_input, value)
			}
		}
	}
	stride := len(clean_input) / len(input)
	fmt.Println(part1(clean_input, stride, row))
}
