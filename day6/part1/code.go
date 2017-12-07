package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var memoryBankFingerprints = make(map[string]bool)

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Could not read the stdin, error %s", err)
		os.Exit(1)
	}

	memorySnapshot := string(raw)
	memoryBank := readMemoryBank(memorySnapshot)
	steps, newMemory := balanceMemoryBank(memoryBank)
	fmt.Printf("Steps: %d\n", steps)
	fmt.Printf("next cycle\n")
	memoryBankFingerprints = make(map[string]bool)
	steps, _ = balanceMemoryBank(newMemory)
	fmt.Printf("Steps: %d\n", steps-1) // n -1, for already seen
}

func readMemoryBank(snapshot string) []int {
	values := strings.Fields(snapshot)
	memoryBank := make([]int, len(values))
	for idx, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic("could not convert value to int")
		}
		memoryBank[idx] = i
	}
	return memoryBank
}

func balanceMemoryBank(memory []int) (int, []int) {
	steps := 0
	for {
		fmt.Printf("memory: %v\n", memory)
		idx := indexForMax(memory)
		blocks := memory[idx]
		memory[idx] = 0

		for i := 1; i <= blocks; i++ {
			pos := (idx + i) % len(memory)
			memory[pos]++
		}
		steps++

		if alreadyBalanced(memory) {
			break
		}
		addMemoryBankFingerprint(memory)
	}
	return steps, memory
}

func addMemoryBankFingerprint(memory []int) {
	h := hashMemoryBank(memory)
	memoryBankFingerprints[h] = true
}

func alreadyBalanced(memory []int) bool {
	h := hashMemoryBank(memory)
	_, ok := memoryBankFingerprints[h]
	return ok
}

func hashMemoryBank(memory []int) string {
	var buf bytes.Buffer
	for idx, v := range memory {
		buf.WriteString(strconv.Itoa(v))
		if idx != len(memory)-1 {
			buf.WriteString("-")
		}
	}
	return buf.String()
}

func indexForMax(ints []int) int {
	idx := 0
	max := 0

	for i, v := range ints {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx
}
