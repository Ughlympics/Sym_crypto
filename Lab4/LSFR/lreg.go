package lsfr

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
	n := 500
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = (L3[i%len(L3)] & L1[i%len(L1)]) ^ ((1 ^ L3[i%len(L3)]) & L2[i%len(L2)])
	}

	return result
}
