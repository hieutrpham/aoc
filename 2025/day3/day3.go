package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(line string) int {
	var max1 byte
	var max2 byte
	for i := 0; i < len(line)-1; i++ {
		if line[i] > max1 {
			max1 = line[i]
			max2 = 0
		}
		for j := i + 1; j < len(line); j++ {
			if line[j] > max2 {
				max2 = line[j]
			}
		}
	}
	return int((max1-'0')*10 + (max2 - '0'))
}

func main() {
	args := os.Args
	path := args[1]
	f, err := os.Open(path)
	check_err(err)
	scanner := bufio.NewScanner(f)
	var num int
	for scanner.Scan() {
		line := scanner.Text()
		num += part1(line)
	}
	fmt.Println(num)
}
