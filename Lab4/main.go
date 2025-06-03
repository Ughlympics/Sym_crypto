package main

import (
	lsfr "Lab4/LSFR"
	"fmt"
)

func main() {
	filename := "var16d.txt"
	count := lsfr.CountBits(filename)
	fmt.Printf("Кількість бітів: %d\n", count)

}

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
