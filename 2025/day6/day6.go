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

func part1(input [][]string, row int) int {
	clean_input := clean(input)
	stride := len(clean_input) / len(input)
	total := 0
	var answer int
	for x := 0; x < stride; x++ {
		op := clean_input[(row-1)*stride+x]
		answer = init_answer(op)
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
	fmt.Println(total)
	return total
}

func clean(input [][]string) []string {
	var clean_input []string
	for row := 0; row < len(input); row++ {
		for item := 0; item < len(input[row]); item++ {
			value := input[row][item]
			if len(value) != 0 {
				clean_input = append(clean_input, value)
			}
		}
	}
	return clean_input
}

func is_empty(str string) bool {
	for i := range str {
		if str[i] != ' ' {
			return false
		}
	}
	return true
}

func init_answer(op string) int {
	if op == "*" {
		return 1
	} else {
		return 0
	}
}

func main() {
	// getting input from stdin
	var input [][]string
	row := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
		row++
	}

	// storing the operator in separate array
	var op []string
	for i := range input[row-1] {
		if !is_empty(input[row-1][i]) {
			op = append(op, input[row-1][i])
		}
	}

	col := len(input[0])
	total := 0
	op_index := 0
	x := 0

	for {
		if op_index >= len(op) {
			break
		}
		answer := init_answer(op[op_index])
		for {
			if x >= col {
				break
			}
			var str string
			for y := 0; y < row-1; y++ {
				str += input[y][x]
			}
			if is_empty(str) {
				x++
				break
			}
			str = strings.Trim(str, " ")
			if op[op_index] == "*" {
				answer *= get_num(str)
			} else {
				answer += get_num(str)
			}
			x++
		}
		op_index++
		total += answer
	}
	fmt.Println("part2 total:", total)
}
