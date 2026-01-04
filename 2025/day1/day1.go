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

func get_dir(s string) rune {
	return rune(s[0])
}

func get_num(s string) int {
	num, err := strconv.Atoi(s[1:])
	check_err(err)
	return num
}

func part1(a []string) int {
	count := 0
	start := 50
	for i := range len(a) {
		if get_dir(a[i]) == 'L' {
			start = start - get_num(a[i])
		} else {
			start = start + get_num(a[i])
		}
		if start%100 == 0 {
			count++
		}
	}
	return count
}

func part2(a []string) int {
	count := 0
	dial := 50
	for i := range len(a) {
		rotation := get_num(a[i])
		if rotation > 100 {
			count += rotation / 100
		}
		rotation = rotation % 100
		if get_dir(a[i]) == 'L' {
			new_dial := dial - rotation
			if new_dial < 0 {
				if dial%100 != 0 {
					count++
				}
				dial = 100 + new_dial
			} else {
				dial = new_dial
			}
		} else if get_dir(a[i]) == 'R' {
			new_dial := dial + rotation
			if new_dial > 100 {
				if dial%100 != 0 {
					count++
				}
				dial = new_dial - 100
			} else {
				dial = new_dial
			}
		}
		if dial%100 == 0 {
			count++
		}
	}
	return count
}

func main() {
	path := "input"
	var a []string
	f, err := os.Open(path)
	check_err(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		check_err(err)
		a = append(a, line)
	}

	// count1 := part1(a)
	// fmt.Println("count part 1:", count1)
	count2 := part2(a)
	fmt.Println("count part 2:", count2)
}
