package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
			result = result*10 + int(i-48)
		} else if i == '-' && result == 0 {
			sign = -1
		}
	}
	return result * sign
}
