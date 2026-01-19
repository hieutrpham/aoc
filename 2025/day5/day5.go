package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func get_id_range(id_range []string) [][]int {
	var s [][]int
	for _, v := range id_range {
		r := strings.Split(v, "-") // array of range: [1, 4]
		r0_atoi, ok := strconv.Atoi(r[0])
		if ok != nil {
			panic("ERR: atoi")
		}
		r1_atoi, ok := strconv.Atoi(r[1])
		if ok != nil {
			panic("ERR: atoi")
		}
		s = append(s, []int{r0_atoi, r1_atoi})
	}
	return s
}

func part1(id_range []string, ingredient_id []string) int {
	sum := 0
	s := get_id_range(id_range)
	for _, v := range ingredient_id {
		for i := range s {
			ret, ok := strconv.Atoi(v)
			if ok != nil {
				continue
			}
			if ret >= s[i][0] && ret <= s[i][1] {
				sum++
				break
			}
		}
	}
	return sum
}

func part2(id_range []string) int {
	s := get_id_range(id_range)
	var sorted_s []int
	for i := range s {
		sorted_s = append(sorted_s, s[i][0])
		sorted_s = append(sorted_s, s[i][1])
	}
	slices.Sort(sorted_s)
	fmt.Println(sorted_s)
	return 0
}

func main() {
	var id_range []string
	var ingredient_id []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		if strings.Contains(line, "-") {
			id_range = append(id_range, line)
		} else {
			ingredient_id = append(ingredient_id, line)
		}
	}
	// fmt.Println(part1(id_range, ingredient_id))
	fmt.Println(part2(id_range))
}
