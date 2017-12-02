package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil || len(input) == 0 {
		panic("Missing values through stdin.")
	}

	strInput := string(input[:])
	fmt.Printf("Code: %s\nResolve:\n%d\n", strInput, resolve(strInput))
}

func resolve(input string) int {
	chars := strings.Split(input, "")

	numbers := make([]int, 0)
	for _, c := range chars {
		if i, err := strconv.Atoi(c); err == nil {
			numbers = append(numbers, i)
		}
	}

	sum := 0
	for i := range numbers {
		n := (i + 1) % len(numbers)
		if numbers[i] == numbers[n] {
			sum += numbers[i]
		}
	}
	return sum
}
