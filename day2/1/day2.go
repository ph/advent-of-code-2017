package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	o := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(o)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		nFields := strings.Fields(scanner.Text())
		numbers := make([]int, 0)
		for _, v := range nFields {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("could not convert to int, error: %s", err))
			}
			numbers = append(numbers, n)
		}

		sort.Ints(numbers)
		sum += numbers[len(numbers)-1] - numbers[0]
	}

	fmt.Printf("Result: %d\n", sum)
}
