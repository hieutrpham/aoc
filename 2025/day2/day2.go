package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(s string) int {
	count := 0
	rangeid := strings.Split(s, "-")
	startid := rangeid[0]
	endid := rangeid[1]
	if len(startid)%2 != 0 && len(endid)%2 != 0 {
		return 0
	}
	startnum, err := strconv.Atoi(rangeid[0])
	check_err(err)
	endnum, err := strconv.Atoi(rangeid[1])
	check_err(err)

	for num := startnum; num <= endnum; num++ {
		s := strconv.Itoa(num)
		middle := len(s) / 2
		for i := range middle {
			if string(s[i:middle]) == string(s[middle:]) {
				count += num
			}
		}
	}
	return count
}

func part2(s string) int {
	count := 0
	rangeid := strings.Split(s, "-")
	startid := rangeid[0]
	endid := rangeid[1]
	startnum, err := strconv.Atoi(startid)
	check_err(err)
	endnum, err := strconv.Atoi(endid)
	check_err(err)

	for num := startnum; num <= endnum; num++ {
		s := strconv.Itoa(num)
		s_len := len(s)
		for i := 1; i <= s_len/2; i++ {
			substr_count := strings.Count(s, s[:i])
			if substr_count >= 2 && i*substr_count == s_len {
				count += num
				break
			}
		}
	}
	return count
}

var count_part_1 = 0
var count_part_2 = 0

func main() {
	args := os.Args
	path := args[1]
	f, err := os.Open(path)
	check_err(err)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	input := strings.Split(line, ",")
	for i := range len(input) {
		count_part_1 += part1(input[i])
		count_part_2 += part2(input[i])
	}
	fmt.Println("count_part_1:", count_part_1)
	fmt.Println("count_part_2:", count_part_2)
}
