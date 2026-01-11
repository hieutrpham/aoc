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

/*
	12 batteries instead of 2

Eg:

	818181911112111
	^ ^ ^ ^^^^^^^^^

create a array of 12 slots. for each slots i have to check for the most optimum number
the first slot can probably be chosen on the first loop.
*/
func part2(line string) int {
	var next int
	var ret [12]int
	for j := 0; j < len(line)-11; j++ {
		if int(line[j]-'0') > ret[0] {
			ret[0] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-10; j++ {
		if int(line[j]-'0') > ret[1] {
			ret[1] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-9; j++ {
		if int(line[j]-'0') > ret[2] {
			ret[2] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-8; j++ {
		if int(line[j]-'0') > ret[3] {
			ret[3] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-7; j++ {
		if int(line[j]-'0') > ret[4] {
			ret[4] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-6; j++ {
		if int(line[j]-'0') > ret[5] {
			ret[5] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-5; j++ {
		if int(line[j]-'0') > ret[6] {
			ret[6] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-4; j++ {
		if int(line[j]-'0') > ret[7] {
			ret[7] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-3; j++ {
		if int(line[j]-'0') > ret[8] {
			ret[8] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-2; j++ {
		if int(line[j]-'0') > ret[9] {
			ret[9] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line)-1; j++ {
		if int(line[j]-'0') > ret[10] {
			ret[10] = int(line[j] - '0')
			next = j + 1
		}
	}
	for j := next; j < len(line); j++ {
		if int(line[j]-'0') > ret[11] {
			ret[11] = int(line[j] - '0')
			next = j + 1
		}
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
