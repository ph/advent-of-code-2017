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

		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		diff := 0
		for i, v1 := range numbers {
			if diff != 0 {
				break
			}
			for _, v2 := range numbers[(i + 1):len(numbers)] {
				if v1%v2 == 0 {
					diff = v1 / v2
					break
				}
			}
		}
		sum += diff
	}

	fmt.Printf("Result: %d\n", sum)
}
