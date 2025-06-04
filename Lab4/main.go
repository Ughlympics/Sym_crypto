package main

import (
	// Пакет з функціями LFSR та Gieffe

	"fmt"
	"math/bits"
	"os"
	"strings"
)

const (
	L1Deg = 25
	L2Deg = 26
	N1Req = 222
	C1    = 71
	N2Req = 229
	C2    = 74
)

func ReadInput(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := strings.TrimSpace(string(data))
	bits := make([]int, len(content))
	for i, ch := range content {
		if ch == '0' {
			bits[i] = 0
		} else {
			bits[i] = 1
		}
	}
	return bits
}

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

func HammingDistance(a, b []int) int {
	count := 0
	for i := range a {
		count += a[i] ^ b[i]
	}
	return count
}

func ReverseBits(val uint32, bits int) uint32 {
	var res uint32
	for i := 0; i < bits; i++ {
		if (val & (1 << i)) != 0 {
			res |= 1 << (bits - 1 - i)
		}
	}
	return res
}

func ConvertSeedsToBitSlices(seeds []uint32, deg uint8) [][]int {
	result := make([][]int, len(seeds))
	for i, seed := range seeds {
		bits := make([]int, deg)
		for j := uint8(0); j < deg; j++ {
			// Отримуємо j-й біт
			if (seed>>j)&1 == 1 {
				bits[deg-j-1] = 1
			} else {
				bits[deg-j-1] = 0
			}
		}
		result[i] = bits
	}
	return result
}

func CheckKeyCompatibility(l1List, l2List []uint32, r uint32) [][2]uint32 {
	var result [][2]uint32

	for _, l1 := range l1List {
		for _, l2 := range l2List {
			compatible := true

			for i := 0; i < 32; i++ {
				rBit := (r >> (31 - i)) & 1
				l1Bit := (l1 >> (31 - i)) & 1
				l2Bit := (l2 >> (31 - i)) & 1

				// Несумісні умови
				if (rBit == 1 && l1Bit == 0 && l2Bit == 0) ||
					(rBit == 0 && l1Bit == 1 && l2Bit == 1) {
					compatible = false
					break
				}
			}

			if compatible {
				result = append(result, [2]uint32{l1, l2})
			}
		}
	}

	return result
}

func generateL1(seed []bool) []bool {
	x := append([]bool(nil), seed...)
	for i := len(seed); i < 32; i++ {
		newBit := x[i-25] != x[i-22] // XOR
		x = append(x, newBit)
	}
	return x
}

// Функція для LFSR L2: y[i+26] = y[i] ⊕ y[i+1] ⊕ y[i+2] ⊕ y[i+6]
func generateL2(seed []bool) []bool {
	y := append([]bool(nil), seed...)
	for i := len(seed); i < 32; i++ {
		newBit := y[i-26] != y[i-25] != y[i-24] != y[i-20]
		y = append(y, newBit)
	}
	return y
}

func ExpandL1Uint32Seeds(seeds []uint32) []uint32 {
	var results []uint32
	for _, seed := range seeds {
		bits := make([]bool, 25)
		for i := 0; i < 25; i++ {
			bits[i] = (seed>>(31-i))&1 == 1
		}
		full := generateL1(bits)
		var val uint32
		for i := 0; i < 32; i++ {
			if full[i] {
				val |= 1 << (31 - i)
			}
		}
		results = append(results, val)
	}
	return results
}

func ExpandL2Uint32Seeds(seeds []uint32) []uint32 {
	var results []uint32
	for _, seed := range seeds {
		bits := make([]bool, 26)
		for i := 0; i < 26; i++ {
			bits[i] = (seed>>(31-i))&1 == 1
		}
		full := generateL2(bits)
		var val uint32
		for i := 0; i < 32; i++ {
			if full[i] {
				val |= 1 << (31 - i)
			}
		}
		results = append(results, val)
	}
	return results
}

func main() {
	z := ReadInput("v11.txt")

	L1Poly := uint32((1 << 3) ^ 1)
	L2Poly := uint32((1 << 6) ^ (1 << 2) ^ (1 << 1) ^ 1)

	var l1Candidates []uint32
	for init := uint32(1); init < (1 << L1Deg); init++ {
		stream := GenerateLFSRStream(L1Poly, L1Deg, init, N1Req)
		R := HammingDistance(stream, z[:N1Req])
		if R <= C1 {
			l1Candidates = append(l1Candidates, init)
		}
	}
	fmt.Printf("Found %d L1 candidates\n", len(l1Candidates))

	var l2Candidates []uint32
	for init := uint32(1); init < (1 << L2Deg); init++ {
		stream := GenerateLFSRStream(L2Poly, L2Deg, init, N2Req)
		R := HammingDistance(stream, z[:N2Req])
		if R <= C2 {
			l2Candidates = append(l2Candidates, init)
		}
	}
	fmt.Printf("Found %d L2 candidates\n", len(l2Candidates))

	fmt.Println("\nFirst 5 L1 candidates (reversed bits):")
	for i := 0; i < len(l1Candidates) && i < 5; i++ {
		fmt.Printf("%025b\n", ReverseBits(l1Candidates[i], L1Deg))
	}

	fmt.Println("\nFirst 5 L2 candidates (reversed bits):")
	for i := 0; i < len(l2Candidates) && i < 5; i++ {
		fmt.Printf("%026b\n", ReverseBits(l2Candidates[i], L2Deg))
	}

	//L3Taps := []int{0, 1, 2, 5}
	//L3Len := 27
	l1Candidats := ExpandL1Uint32Seeds(l1Candidates)
	l2Candidats := ExpandL2Uint32Seeds(l2Candidates)
	r := uint32(0b1111101001010000111000111110) // тестовий приклад

	compatiblePairs := CheckKeyCompatibility(l1Candidats, l2Candidats, r)

	fmt.Printf("Знайдено %d сумісних пар:\n", len(compatiblePairs))
	for _, pair := range compatiblePairs {
		fmt.Println("L1:", pair[0])
		fmt.Println("L2:", pair[1])
		fmt.Println()
	}

}

// func ReadInputBits(filename string) ([]int, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var bits []int
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for _, ch := range line {
// 			if ch == '0' || ch == '1' {
// 				bits = append(bits, int(ch-'0'))
// 			}
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return bits, nil
// }

// // Генерація LFSR-потоку
// func GenerateLFSRStream(poly uint32, deg uint8, init uint32, steps int) []int {
// 	state := init
// 	result := make([]int, steps)
// 	for i := 0; i < steps; i++ {
// 		result[i] = int(state & 1)
// 		newBit := bits.OnesCount32(state&poly) % 2
// 		state = (state >> 1) | (uint32(newBit) << (deg - 1))
// 	}
// 	return result
// }

// // Функція обчислення Hamming-відстані
// func HammingDistance(a, b []int) int {
// 	count := 0
// 	for i := range a {
// 		count += a[i] ^ b[i]
// 	}
// 	return count
// }

// func FindCandidates(poly uint32, deg uint8, N_req int, C int, z []int) []uint32 {
// 	var candidates []uint32
// 	for init := uint32(1); init < (1 << deg); init++ {
// 		stream := GenerateLFSRStream(poly, deg, init, N_req)
// 		R := HammingDistance(stream, z[:N_req])
// 		if R <= C {
// 			candidates = append(candidates, init)
// 		}
// 	}
// 	return candidates
// }

// func main() {
// 	// Параметри
// 	L1_poly := uint32((1 << 3) ^ 1) // x^3 + 1
// 	L1_deg := uint8(25)
// 	N1_req := 222
// 	C1 := 71

// 	L2_poly := uint32((1 << 6) ^ (1 << 2) ^ (1 << 1) ^ 1) // x^6 + x^2 + x + 1
// 	L2_deg := uint8(26)
// 	N2_req := 229
// 	C2 := 73

// 	// Зчитування z
// 	z, err := ReadInputBits("v11.txt")
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	// Кандидати на L1
// 	fmt.Println("Пошук кандидатів для L1...")
// 	l1Candidates := FindCandidates(L1_poly, L1_deg, N1_req, C1, z)
// 	fmt.Printf("Знайдено %d кандидатів для L1\n", len(l1Candidates))

// 	// Кандидати на L2
// 	fmt.Println("Пошук кандидатів для L2...")
// 	l2Candidates := FindCandidates(L2_poly, L2_deg, N2_req, C2, z)
// 	fmt.Printf("Знайдено %d кандидатів для L2\n", len(l2Candidates))

// 	// За бажанням: роздрукувати знайдені значення у двійковому вигляді
// 	fmt.Println("\nПерші 5 кандидатів L1:")
// 	for i := 0; i < len(l1Candidates) && i < 186; i++ {
// 		fmt.Printf("%025b\n", l1Candidates[i])
// 	}

// 	fmt.Println("\nПерші 5 кандидатів L2:")
// 	for i := 0; i < len(l2Candidates) && i < 5; i++ {
// 		fmt.Printf("%026b\n", l2Candidates[i])
// 	}
// }
