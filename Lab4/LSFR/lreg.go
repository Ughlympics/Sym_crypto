package lsfr

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LFSR(seed []int, taps []int, steps int) []int {
	state := append([]int(nil), seed...)
	result := make([]int, 0, steps)

	for i := 0; i < steps; i++ {
		result = append(result, state[0])

		newBit := 0
		for _, t := range taps {
			newBit ^= state[t]
		}

		state = append(state[1:], newBit)
	}
	return result
}

func GieffeGenerator(L1, L2, L3 []int) []int {
	n := 2048
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = (L3[i%len(L3)] & L1[i%len(L1)]) ^ ((1 ^ L3[i%len(L3)]) & L2[i%len(L2)])
	}

	return result
}

// ReadInput loads the binary string from the file
func ReadInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bits []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			if ch == '0' || ch == '1' {
				bit, _ := strconv.Atoi(string(ch))
				bits = append(bits, bit)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bits
}

////////////////

func FindCandidates(inputSequence []int, lfsrLen int, taps []int) []string {
	var candidates []string
	total := 1 << lfsrLen // 2^length

	for i := 0; i < total; i++ {
		seed := intToBits(i, lfsrLen)
		output := LFSR(seed, taps, len(inputSequence))

		match := true
		for j := 0; j < len(inputSequence); j++ {
			if output[j] != inputSequence[j] {
				match = false
				break
			}
		}

		if match {
			candidates = append(candidates, bitsToString(seed))
		}
	}
	return candidates
}

func intToBits(n int, length int) []int {
	bits := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		bits[i] = (n >> (length - 1 - i)) & 1
	}
	return bits
}

func bitsToString(bits []int) string {
	var s string
	for _, b := range bits {
		s += strconv.Itoa(b)
	}
	return s
}
