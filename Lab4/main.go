package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func ReadInputBits(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bits []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			if ch == '0' || ch == '1' {
				bits = append(bits, int(ch-'0'))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return bits, nil
}

// Генерація LFSR-потоку
func GenerateLFSRStream(poly uint32, deg uint8, init uint32, steps int) []int {
	state := init
	result := make([]int, steps)
	for i := 0; i < steps; i++ {
		result[i] = int(state & 1)
		newBit := bits.OnesCount32(state&poly) % 2
		state = (state >> 1) | (uint32(newBit) << (deg - 1))
	}
	return result
}

// Функція обчислення Hamming-відстані
func HammingDistance(a, b []int) int {
	count := 0
	for i := range a {
		count += a[i] ^ b[i]
	}
	return count
}

func FindCandidates(poly uint32, deg uint8, N_req int, C int, z []int) []uint32 {
	var candidates []uint32
	for init := uint32(1); init < (1 << deg); init++ {
		stream := GenerateLFSRStream(poly, deg, init, N_req)
		R := HammingDistance(stream, z[:N_req])
		if R <= C {
			candidates = append(candidates, init)
		}
	}
	return candidates
}

func main() {
	// Параметри
	L1_poly := uint32((1 << 3) ^ 1) // x^3 + 1
	L1_deg := uint8(25)
	N1_req := 222
	C1 := 71

	L2_poly := uint32((1 << 6) ^ (1 << 2) ^ (1 << 1) ^ 1) // x^6 + x^2 + x + 1
	L2_deg := uint8(26)
	N2_req := 229
	C2 := 74

	// Зчитування z
	z, err := ReadInputBits("v.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Кандидати на L1
	fmt.Println("Пошук кандидатів для L1...")
	l1Candidates := FindCandidates(L1_poly, L1_deg, N1_req, C1, z)
	fmt.Printf("Знайдено %d кандидатів для L1\n", len(l1Candidates))

	// Кандидати на L2
	fmt.Println("Пошук кандидатів для L2...")
	l2Candidates := FindCandidates(L2_poly, L2_deg, N2_req, C2, z)
	fmt.Printf("Знайдено %d кандидатів для L2\n", len(l2Candidates))

	// За бажанням: роздрукувати знайдені значення у двійковому вигляді
	fmt.Println("\nПерші 5 кандидатів L1:")
	for i := 0; i < len(l1Candidates) && i < 5; i++ {
		fmt.Printf("%025b\n", l1Candidates[i])
	}

	fmt.Println("\nПерші 5 кандидатів L2:")
	for i := 0; i < len(l2Candidates) && i < 5; i++ {
		fmt.Printf("%026b\n", l2Candidates[i])
	}
}
