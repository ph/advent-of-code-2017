package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		hashedWord := hash(word)
		v, ok := count[hashedWord]
		if ok && v == 1 {
			valid = false
			break
		}
		count[hashedWord]++
	}
	return valid
}

func hash(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}
