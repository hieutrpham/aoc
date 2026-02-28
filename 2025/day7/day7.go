package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print_array[T any](input []T) {
	for i := range input {
		fmt.Println(input[i])
	}
}

func part1(input [][]string) ([][]string, int) {
	row := len(input)
	col := len(input[0])
	count := 0
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			curr_char := input[y][x]
			if curr_char == "S" {
				input[y+1][x] = "|"
			}
			if curr_char == "^" && input[y-1][x] == "|" {
				input[y+1][x-1] = "|"
				input[y+1][x+1] = "|"
				count++
			}
			if curr_char == "." && y > 0 && input[y-1][x] == "|" {
				input[y][x] = "|"
			}
		}
	}
	return input, count
}

func get_input() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	var input [][]string
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}
	return input
}

/*
after each splitter, there are only 2 paths, left or right that a particle can go down

		  |
		  ^
		 1 1 -> 1 way to get to each timeline. sum = 2
		 ^ ^
		1 2 1 -> 1 way to get to rear timeline. 2 ways in the middle one. sum = 4
		^ ^ ^
	   1 3 3 1
	   ^ ^ _ ^
	  1 4 331 1 -> if there is no previous splitter, don't add the 2 previous timelines
*/
/*
how to build a binary tree?

	a beam/pipe
	a spliter?
		a splitter denotes a new node if there is a beam on top
	how to connect the nodes to create the tree?
		one node could be connected to the node before by 1 column to the left or right and 2 rows up
		each beam will have information of the node that created it
		which means i have to create a struct of beam that store the node parent
		the beam under each beam will have the same information as the one before it
	looping through the input and construct the tree line by line
		in order to know which node connects to which, i have to take a look back to the previous coord
		that means when i create a node on row 2, i don't know how to connect the left and right node yet?
		only when i move to row 4 that i can connect the node in row 2
		add to the left node if based on the col parameter
*/

// coordinate of each node
type Coord struct {
	row int
	col int
}

type Node struct {
	value    int
	coord    Coord
	children []*Node
}

func newNode(row int, col int) Node {
	return Node{coord: Coord{row, col}}
}

// if there are no children, the node value will be 1
// if there are children, the node value will be the sum of the children's node values
func count_timeline(node *Node) int {
	if node == nil {
		return 1
	}
	if node != nil && len(node.children) == 0 {
		node.value += 1
	}
	if node != nil && len(node.children) > 0 {
		for _, child := range node.children {
			if child != nil && child.value != 0 {
				node.value += child.value
			} else {
				node.value += count_timeline(child)
			}
		}
	}
	return node.value
}

func main() {
	input := get_input()
	output, _ := part1(input)

	// root node
	nodes := make(map[Coord]*Node)
	row_len := len(output)
	col_len := len(output[0])
	var start_pos Coord
	// create all the beam structs
	for row := 0; row < row_len; row++ {
		for col := 0; col < col_len; col++ {
			if output[row][col] == "|" {
				if output[row-1][col] == "S" {
					start_pos = Coord{row, col}
				}
				if nodes[Coord{row, col}] == nil {
					new := newNode(row, col)
					nodes[Coord{row, col}] = &new
				}
			}
		}
	}
	// connect the beams
	for row := 0; row < row_len; row++ {
		for col := 0; col < col_len; col++ {
			if output[row][col] == "|" {
				if row < row_len-1 && output[row+1][col] == "^" {
					nodes[Coord{row, col}].children = append(nodes[Coord{row, col}].children, nodes[Coord{row + 2, col - 1}])
					nodes[Coord{row, col}].children = append(nodes[Coord{row, col}].children, nodes[Coord{row + 2, col + 1}])
				} else if row < row_len-1 && output[row+1][col] == "|" {
					nodes[Coord{row, col}].children = append(nodes[Coord{row, col}].children, nodes[Coord{row + 1, col}])
				}
			}
		}
	}
	fmt.Println("total timeline:", count_timeline(nodes[start_pos]))
}
