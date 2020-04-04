//Test CI Flow

package main

import (
	"fmt"
	"strings"
)

func findDirections(p string) int {
	startingFloor := 0
	floorsInput := strings.Split(p, "")
	for _, v := range floorsInput {
		if v == "(" {
			startingFloor++
		} else if v == ")" {
			startingFloor--
		}
	}
	return startingFloor
}

func main() {
	direction := findDirections("()(")
	fmt.Println(direction)
}
