package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Could not read stdin, err: %s\n", err)
	}
	var instructions []int
	for _, v := range strings.Split(string(raw), "\n") {
		if v == "" {
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Could not convert: `%s`, err: %v\n", v, err)
		}
		instructions = append(instructions, i)
	}

	programCounter := 0
	steps := 0

	for programCounter < len(instructions) {
		instruction := instructions[programCounter]
		if instruction >= 3 {
			instructions[programCounter]--
		} else {
			instructions[programCounter]++
		}
		programCounter += instruction
		steps++
	}

	fmt.Printf("Steps: %d\n", steps)
}
