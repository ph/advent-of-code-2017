package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	o := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(o)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		if validPassword(scanner.Text()) {
			sum++
		}
	}
	fmt.Printf("Response: %d\n", sum)
}

func validPassword(p string) bool {
	valid := true
	words := strings.Fields(p)
	count := make(map[string]int, 0)
	for _, word := range words {
		v, ok := count[word]
		if ok && v == 1 {
			valid = false
			break
		}
		count[word]++
	}
	return valid
}
