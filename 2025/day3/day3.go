package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part2(line string) int {
	var next int
	var ret [12]int

	// this function fill each slot of the ret array with the optimum digit
	// once it finds the optimum number for each slot, it moves on to the 'next' index
	fill_slot := func(line string, slot int) {
		for j := next; j < len(line)-(11-slot); j++ {
			if int(line[j]-'0') > ret[slot] {
				ret[slot] = int(line[j] - '0')
				next = j + 1
			}
		}
	}

	// just call the fill slot function 12 times to fill all 12 batteries
	for i := 0; i < 12; i++ {
		fill_slot(line, i)
	}

	var s string
	for i := 0; i < len(ret); i++ {
		s += string(ret[i] + '0')
	}
	num, err := strconv.Atoi(s)
	check_err(err)
	return num
}

func main() {
	args := os.Args
	path := args[1]
	f, err := os.Open(path)
	check_err(err)
	scanner := bufio.NewScanner(f)
	var part1_num int
	var part2_num int
	for scanner.Scan() {
		line := scanner.Text()
		part1_num += part1(line)
		part2_num += part2(line)
	}
	// fmt.Println("part 1:", part1_num)
	fmt.Println("part 2:", part2_num)
}
