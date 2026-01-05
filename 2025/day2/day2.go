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

func get_valid_id(s string) int {
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

var count = 0

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
		count += get_valid_id(input[i])
	}
	fmt.Println("count:", count)
}
