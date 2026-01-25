package main

import (
	"bufio"
	"cmp"
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

const max = 1
const min = 0

// checking if given a range of id []int overlap with the existing ranges [][]int
// if detect overlap, return true
func overlap(id_ranges [][]int, value_range []int) bool {
	for i := range id_ranges {
		// fmt.Println("LOG:", value_range, id_ranges[i])
		if value_range[min] == id_ranges[i][min] && value_range[max] == id_ranges[i][max] {
			return true
		}
		if !(value_range[min] > id_ranges[i][max] || value_range[max] < id_ranges[i][min]) {
			return true
		}
	}
	return false
}

// introducing the struct to hold ingredient id ranges and a boolean field to signal if it was counted when summing the ranges
type ingredient struct {
	id_range []int
	included bool
}

func part2(id_range []string) int {
	// array of ints converted from array of string
	s := get_id_range(id_range)
	slices.SortFunc(s, func(a, b []int) int {
		return cmp.Compare(a[min], b[min])
	})
	// fmt.Println(s)
	// array to store the overlap id ranges
	var overlap_arr []ingredient
	var rest_arr []ingredient

	// for each range in s, loop through every item in s and check if there are overlap
	for i := range s {
		new := ingredient{s[i], false}
		if overlap(s, s[i]) {
			overlap_arr = append(overlap_arr, new)
		} else {
			rest_arr = append(rest_arr, new)
		}
	}
	// fmt.Println(overlap_arr)
	sum := 0
	l_overlap := len(overlap_arr)
	// for each range in overlap, check for every other ranges to see which ones overlapping each other, then summing the max - min + 1
	for i := 0; i < l_overlap; i++ {
		if !overlap_arr[i].included {

			var s_arr [][]int
			s_arr = append(s_arr, overlap_arr[i].id_range)

			for j := 0; j < l_overlap; j++ {
				if !overlap_arr[j].included && overlap(s_arr, overlap_arr[j].id_range) {
					s_arr = append(s_arr, overlap_arr[j].id_range)
					overlap_arr[j].included = true
				}
			}

			var ret_arr []int
			for i := range s_arr {
				ret_arr = append(ret_arr, s_arr[i][min])
				ret_arr = append(ret_arr, s_arr[i][max])
			}

			sum += slices.Max(ret_arr) - slices.Min(ret_arr) + 1
			overlap_arr[i].included = true
		}
	}

	// summing the non overlapping
	for i := range rest_arr {
		sum += rest_arr[i].id_range[max] - rest_arr[i].id_range[min] + 1
	}

	fmt.Println("sum:", sum)
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
	part2(id_range)
}
