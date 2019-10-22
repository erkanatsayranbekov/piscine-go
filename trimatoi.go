package piscine

func TrimAtoi(s string) int {
	array := []rune{}
	result := 0
	sign := 1
	error := 0
	for _, i := range s {
		if i == '-' && error == 0 {
			sign = -1
		} else if i >= '0' && i <= '9' {
			array = append(array, i)
			error++
		} else {
			continue
		}
	}

	for _, i := range array {
		num := 0
		for j := '1'; j <= i; j++ {

			num++
		}
		result = result*10 + num
	}
	return result * sign
}
