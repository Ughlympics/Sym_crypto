package main

import (
	lsfr "Lab4/LSFR"
	"fmt"
	"math"
	"math/bits"
	"os"
)

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

func main() {
	L1_poly := uint32((1 << 3) ^ 1)
	L1_deg := uint8(25)
	L2_poly := uint32((1 << 6) ^ (1 << 2) ^ (1 << 1) ^ 1)
	L2_deg := uint8(26)
	L3_poly := uint32((1 << 5) ^ (1 << 2) ^ (1 << 1) ^ 1)
	L3_deg := uint8(27)

	r_seq := lsfr.ReadInput("v.txt")
	N := len(r_seq)
	N1_req := 222
	//C1 := 71
	N2_req := 229
	//C2 := 74
	//maxOffset := 64

	// Пошук найкращого кандидата для L1
	fmt.Println("Start best L1 candidate search")
	var bestL1 uint32
	minDev := float64(N1_req)
	for init := uint32(1); init < (1 << L1_deg); init++ {
		stream := GenerateLFSRStream(L1_poly, L1_deg, init, N1_req)
		R := 0
		for i := 0; i < N1_req; i++ {
			R += stream[i] ^ r_seq[i]
		}
		dev := math.Abs(float64(R) - 0.25*float64(N1_req))
		if dev < minDev {
			minDev = dev
			bestL1 = init
		}
	}
	fmt.Printf("Best L1: %025b (dev=%.2f)\n", bestL1, minDev)

	// Пошук найкращого кандидата для L2
	fmt.Println("Start best L2 candidate search")
	var bestL2 uint32
	minDev = float64(N2_req)
	for init := uint32(1); init < (1 << L2_deg); init++ {
		stream := GenerateLFSRStream(L2_poly, L2_deg, init, N2_req)
		R := 0
		for i := 0; i < N2_req; i++ {
			R += stream[i] ^ r_seq[i]
		}
		dev := math.Abs(float64(R) - 0.25*float64(N2_req))
		if dev < minDev {
			minDev = dev
			bestL2 = init
		}
	}
	fmt.Printf("Best L2: %026b (dev=%.2f)\n", bestL2, minDev)

	// Генеруємо повну послідовність L3
	//fullLen := (1 << L3_deg) + N
	currCandidate := uint32(1)
	//L3_full := GenerateLFSRStream(L3_poly, L3_deg, currCandidate, fullLen)

	// Генеруємо послідовності L1 та L2
	L1_stream := GenerateLFSRStream(L1_poly, L1_deg, bestL1, N)
	L2_stream := GenerateLFSRStream(L2_poly, L2_deg, bestL2, N)

	// Пошук зсуву
	fmt.Println("Start L3 candidate search")

	for l3_init := uint32(1); l3_init < (1 << L3_deg); l3_init++ {
		L3_stream := GenerateLFSRStream(L3_poly, L3_deg, l3_init, N)

		match := true
		for i := 0; i < N; i++ {
			out := (L3_stream[i] & L1_stream[i]) ^ ((1 ^ L3_stream[i]) & L2_stream[i])
			if out != r_seq[i] {
				match = false
				break
			}
		}

		if match {
			fmt.Printf("Match found!\nL1: %025b\nL2: %026b\nL3 init: %027b\n", bestL1, bestL2, l3_init)
			os.Exit(0)
		}

		// Обчислюємо наступний стан для L3 (аналог зсува в C++)
		nextBit := bits.OnesCount32(currCandidate&L3_poly) % 2
		currCandidate = (currCandidate >> 1) | (uint32(nextBit) << (L3_deg - 1))
	}

	fmt.Println("No match found")
}

/////...../////..../////....////.....////

// func main() {
// 	// Параметри
// 	L1_poly := uint32((1 << 3) ^ 1)
// 	L1_deg := uint8(25)
// 	L2_poly := uint32((1 << 6) ^ (1 << 2) ^ (1 << 1) ^ 1)
// 	L2_deg := uint8(26)
// 	L3_poly := uint32((1 << 5) ^ (1 << 2) ^ (1 << 1) ^ 1)
// 	L3_deg := uint8(27)

// 	r_seq := lsfr.ReadInput("v.txt")
// 	N := len(r_seq)
// 	N1_req := 222
// 	C1 := 71
// 	N2_req := 229
// 	C2 := 74
// 	maxOffset := 64 // <= Обмеження кількості зсувів

// 	fmt.Println("Start L1 candidate search")
// 	L1_candidates := make([]uint32, 0)
// 	for init := uint32(1); init < (1 << L1_deg); init++ {
// 		stream := GenerateLFSRStream(L1_poly, L1_deg, init, int(N1_req))
// 		R := 0
// 		for i := 0; i < N1_req; i++ {
// 			R += stream[i] ^ r_seq[i]
// 		}
// 		if R < C1 {
// 			L1_candidates = append(L1_candidates, init)
// 		}
// 	}
// 	fmt.Println("L1 candidates:", len(L1_candidates))

// 	fmt.Println("Start L2 candidate search")
// 	L2_candidates := make([]uint32, 0)
// 	for init := uint32(1); init < (1 << L2_deg); init++ {
// 		stream := GenerateLFSRStream(L2_poly, L2_deg, init, int(N2_req))
// 		R := 0
// 		for i := 0; i < N2_req; i++ {
// 			R += stream[i] ^ r_seq[i]
// 		}
// 		if R < C2 {
// 			L2_candidates = append(L2_candidates, init)
// 		}
// 	}
// 	fmt.Println("L2 candidates:", len(L2_candidates))

// 	fmt.Println("Start L3 search")
// 	for _, l1 := range L1_candidates {
// 		L1_stream := GenerateLFSRStream(L1_poly, L1_deg, l1, N)
// 		for _, l2 := range L2_candidates {
// 			L2_stream := GenerateLFSRStream(L2_poly, L2_deg, l2, N)
// 			for l3_init := uint32(1); l3_init < (1 << L3_deg); l3_init++ {
// 				L3_full := GenerateLFSRStream(L3_poly, L3_deg, l3_init, N+maxOffset)

// 				for offset := 0; offset < maxOffset; offset++ {
// 					match := true
// 					for i := 0; i < N; i++ {
// 						l3bit := L3_full[offset+i]
// 						out := (l3bit & L1_stream[i]) ^ ((1 ^ l3bit) & L2_stream[i])
// 						if out != r_seq[i] {
// 							match = false
// 							break
// 						}
// 					}
// 					if match {
// 						fmt.Printf("Found match!\nL1: %025b\nL2: %026b\nL3: %027b offset: %d\n", l1, l2, l3_init, offset)
// 						os.Exit(0)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("No match found")
// }

// func main() {
// 	L1Seed := make([]int, 30)
// 	for i := range L1Seed {
// 		L1Seed[i] = 1
// 	}
// 	L2Seed := make([]int, 31)
// 	for i := range L2Seed {
// 		L2Seed[i] = 1
// 	}
// 	L3Seed := make([]int, 32)
// 	for i := range L3Seed {
// 		L3Seed[i] = 1
// 	}

// 	L1Taps := []int{0, 1, 4, 6}       // x^0, x^1, x^4, x^6
// 	L2Taps := []int{0, 3}             // x^0, x^3
// 	L3Taps := []int{0, 1, 2, 3, 5, 7} // x^0, x^1, x^2, x^3, x^5, x^7

// 	// Генеруємо
// 	gg1 := make([]int, 30)
// 	gg1 = lsfr.LFSR(L1Seed, L1Taps, 500)

// 	gg2 := make([]int, 31)
// 	gg2 = lsfr.LFSR(L2Seed, L2Taps, 500)

// 	gg3 := make([]int, 32)
// 	gg3 = lsfr.LFSR(L3Seed, L3Taps, 500)

// 	gg := make([]int, 30)
// 	gg = lsfr.GieffeGenerator(gg1, gg2, gg3)

// 	fmt.Println("Gieffe Generator Results:", gg)

// 	fmt.Println("L1:", lsfr.LFSR(L1Seed, L1Taps, 500))
// 	fmt.Println("L2:", lsfr.LFSR(L2Seed, L2Taps, 50))
// 	fmt.Println("L3:", lsfr.LFSR(L3Seed, L3Taps, 50))
// }
