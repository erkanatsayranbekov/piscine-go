package piscine

func StrLen(str string) int {
	a := 0
	s := []rune(str)
	for i := range s {
		a = i
	}
	return a
}

var m = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	var di int
	for _, digit := range s {
		if d, ok := m[digit]; ok {
			di = (di * 10) + d
		} else {
			return 0
		}
	}
	return di * sign
}
