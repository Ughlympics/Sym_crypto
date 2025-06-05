package lsfr

//реалізація генератора Джиффі на булеві значення
//Полегшуує фунцію порівння коли генеруємо послідовність на 200 біт

const checkBits = 200

func GenerateL1(seed []bool) []bool {
	x := append([]bool(nil), seed...)
	for i := len(seed); i < checkBits; i++ {
		newBit := x[i-25] != x[i-22]
		x = append(x, newBit)
	}
	return x
}

func GenerateL2(seed []bool) []bool {
	y := append([]bool(nil), seed...)
	for i := len(seed); i < checkBits; i++ {
		newBit := y[i-26] != y[i-25] != y[i-24] != y[i-20]
		y = append(y, newBit)
	}
	return y
}

func GenerateL3(seed []bool) []bool {
	s := append([]bool(nil), seed...)
	for i := len(seed); i < checkBits; i++ {
		newBit := s[i-27] != s[i-26] != s[i-25] != s[i-22]
		s = append(s, newBit)
	}
	return s
}

func F(x, y, s bool) bool {
	if s {
		return x
	}
	return y
}

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
