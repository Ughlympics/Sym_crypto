package main

import (
	"Lab4/lsfr"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

const (
	L1Deg     = 25
	L2Deg     = 26
	N1Req     = 222
	C1        = 71
	N2Req     = 229
	C2        = 74
	checkBits = 200
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

			for i := 0; i < 25; i++ {
				rBit := (r >> (24 - i)) & 1

				var l1Bit, l2Bit uint32

				if i < 25 {
					l1Bit = (l1 >> (24 - i)) & 1
				}
				if i < 26 {
					l2Bit = (l2 >> (25 - i)) & 1
				}

				// Несумісність: якщо r = 1, а обидва 0 — не годиться;
				// або r = 0, а обидва 1 — теж не годиться.
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

// ...///

func generateL1FromSeed(seed uint32, length int) []bool {
	x := make([]bool, 0, length)
	for i := 0; i < 25; i++ {
		x = append(x, (seed>>(24-i))&1 == 1)
	}
	for i := 25; i < length; i++ {
		newBit := x[i-25] != x[i-22] // XOR
		x = append(x, newBit)
	}
	return x
}

func generateL2FromSeed(seed uint32, length int) []bool {
	y := make([]bool, 0, length)
	for i := 0; i < 26; i++ {
		y = append(y, (seed>>(25-i))&1 == 1)
	}
	for i := 26; i < length; i++ {
		newBit := y[i-26] != y[i-25] != y[i-24] != y[i-20]
		y = append(y, newBit)
	}
	return y
}

func CheckKeyPairCompatibility(candidates [][2]uint32, input []bool) [][2]uint32 {
	var result [][2]uint32

	for _, pair := range candidates {
		l1, l2 := pair[0], pair[1]

		x := generateL1FromSeed(l1, checkBits)
		y := generateL2FromSeed(l2, checkBits)

		compatible := true
		for i := 0; i < checkBits; i++ {
			rBit := input[i]
			xBit := x[i]
			yBit := y[i]

			if (rBit && !xBit && !yBit) || (!rBit && xBit && yBit) {
				compatible = false
				break
			}
		}

		if compatible {
			result = append(result, pair)
		}
	}

	return result
}

func stringToBoolSlice(s string) []bool {
	result := make([]bool, 0, len(s))
	for _, c := range s {
		if c == '1' {
			result = append(result, true)
		} else if c == '0' {
			result = append(result, false)
		}

	}
	return result
}

//...///

func FindL3(l1Seed, l2Seed, input []bool) []bool {
	x := lsfr.GenerateL1(l1Seed)
	y := lsfr.GenerateL2(l2Seed)

	known := make([]int, checkBits)
	for i := range known {
		known[i] = -1
		if x[i] == true && y[i] == false {
			if input[i] == true {
				known[i] = 1
			} else {
				known[i] = 0
			}
		} else if x[i] == true && y[i] == false {
			if input[i] == true {
				known[i] = 0
			} else {
				known[i] = 1
			}
		}
	}

	for candidate := 0; candidate < (1 << 27); candidate++ {
		seed := make([]bool, 27)
		for i := 0; i < 27; i++ {
			seed[26-i] = (candidate>>i)&1 == 1
		}

		s := lsfr.GenerateL3(seed)

		match := true
		for i := 0; i < checkBits; i++ {
			if known[i] != -1 && lsfr.B2i(s[i]) != known[i] {
				match = false
				break
			}

			if input[i] != lsfr.F(x[i], y[i], s[i]) {
				match = false
				break
			}
		}

		if match {
			fmt.Println("Знайдено L3:")
			return seed
		}
	}

	fmt.Println("Не знайдено L3.")
	return nil
}

// ////////////////
// ///////////////
func CheckKeyCompatibility2(l1List, l2List []uint32, r uint32) [][2]uint32 {
	var result [][2]uint32

	for _, l1 := range l1List {
		for _, l2 := range l2List {
			compatible := true

			for i := 0; i < 25; i++ {
				rBit := (r >> (24 - i)) & 1
				l1Bit := (l1 >> (24 - i)) & 1
				l2Bit := (l2 >> (25 - i)) & 1

				// Несумісність:
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

///////////////
///////////////

func ReverseCandidates(candidates []uint32, bits int) []uint32 {
	reversed := make([]uint32, len(candidates))
	for i, val := range candidates {
		reversed[i] = ReverseBits(val, bits)
	}
	return reversed
}

func uintToBoolSlice(val uint32, size int) []bool {
	bits := make([]bool, size)
	for i := 0; i < size; i++ {
		bits[i] = ((val >> (size - 1 - i)) & 1) == 1
	}
	return bits
}

func main() {
	z := ReadInput("var16d.txt")
	//ciphertext := stringToBoolSlice("11111010010100001110001111101110010001111110101000011011000010001001000001100111111110001110100111110100011101111010011110100000001000001110000000011000000000011100011111101110100100110111101100101100011111110000001101010110110000110001111111110011011010011010000110111011101111011001101000111000111001100100010100110111001110001100011100100011011101100110011010001110001010101110000001011111000110101011010101000001001010101110011010101101000110011100110111001010101001010111000110101001000011001110111011110000001011011110111000110010100110110101010000011011011101001011100010001101000101001011110001011010111001101010111001001010110101010100001010111000110101110101111111111101001111011101100100101100000010010010000011010001011101100101110011111100100000111110100110000101010011111001100101010011000101001000101010100101110110101011001100011000010101110011000011111100100111011111111010101011011100101010100101000100111101001010001010110001001010000000110101101110001101110110001001110010010100101110000001110000011100001111110110001111011110011101101001100100101101111011010101001011100000010100111011000000011100011101110000101010100000011011111110011111110010000010010100100000111100001001111000111011001101101101011001100011000111101100010101011100110111000110011101100001101110010011111100000011101000000001110000100010101100100111011110011110101001110101011110000111101000110100010100000101010110001010111011010111111110101000100100111111111101000110010011101101001110110000010100110000101001010101010100110111111011011101001101010101001010011000100100110111001101101001001101010011001100110001010011001100111001110000100010001011110011101111011100110101001110001110000011101100100101100100010100010101011001011100000101001111001010100000001101011111101101000110000011110101101000011010110010100100111000101000001101111101001010011101111111101101100111001111110110110110111010100011011111101000110110010100011110000111010011011111000000111010001001001100000111010011100110011110101110000111110101010010100101100001010001000110010000110110")

	//Беремо 200 і 25 перших біт з нашого варіанту
	//var11
	//cipher200 := stringToBoolSlice("11111010010100001110001111101110010001111110101000011011000010001001000001100111111110001110100111110100011101111010011110100000001000001110000000011000000000011100011111101110100100110111101100101100") // 200 біт
	//first25 := uint32(0b1111101001010000111000111)

	//var16
	cipher200 := stringToBoolSlice("10001001110110011100110110100110010100010010011101110011111100000100001111010011010001001000111001001101110110000111111011010000100000101110111111100110111101011101001110101011010101011010011000100110")
	first25 := uint32(0b1000100111011001110011011)

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
	fmt.Printf("Знайдено %d L1 кандидатів\n", len(l1Candidates))

	var l2Candidates []uint32
	for init := uint32(1); init < (1 << L2Deg); init++ {
		stream := GenerateLFSRStream(L2Poly, L2Deg, init, N2Req)
		R := HammingDistance(stream, z[:N2Req])
		if R <= C2 {
			l2Candidates = append(l2Candidates, init)
		}
	}
	fmt.Printf("Знайдено %d L2 кандидатів\n", len(l2Candidates))

	l1Reversed := ReverseCandidates(l1Candidates, 25)
	l2Reversed := ReverseCandidates(l2Candidates, 26)

	compatiblePairs := CheckKeyCompatibility2(l1Reversed, l2Reversed, first25)

	fmt.Printf("Знайдено %d сумісних пар\n", len(compatiblePairs))

	//робимо пошук догенеровуючи до довжини 200 біт L1 та L2
	valid := CheckKeyPairCompatibility(compatiblePairs, cipher200)

	for _, pair := range valid {
		fmt.Printf("Потрібний L1: %025b\n", pair[0])
		fmt.Printf("Потрібний L2: %026b\n", pair[1])
	}

	l1 := make([]bool, 0, 25)
	l2 := make([]bool, 0, 26)

	if len(valid) > 0 {
		l1 = uintToBoolSlice(valid[0][0], 25)
		l2 = uintToBoolSlice(valid[0][1], 26)
	}

	//l1 := stringToBoolSlice("1110101001011000111101101")                                                                                                                                                                                    // 25 бит
	//l2 := stringToBoolSlice("01111111000101111110101111")                                                                                                                                                                                   // 26 бит

	l3 := FindL3(l1, l2, cipher200)
	if l3 != nil {
		fmt.Print("L3: ")
		for _, b := range l3 {
			fmt.Print(lsfr.B2i(b))
		}
		fmt.Println()
	}

}
