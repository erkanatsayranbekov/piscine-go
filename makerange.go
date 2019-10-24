package piscine

func MakeRange(min, max int) []int {

	if min < max {
		answer := make([]int, max-min)
		j := 0
		for i := min; i < max; i++ {
			answer[j] = i
			j++
		}
		return answer
	}
	nil := make([]int, 0)
	return nil
}
