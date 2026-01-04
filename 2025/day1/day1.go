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

func main() {
	path := "./day1/input"
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
	fmt.Println("count:", count)
}
