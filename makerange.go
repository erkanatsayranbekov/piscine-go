package piscine

func MakeRange(min, max int) []int {

	if min < max && min > 0 && max > 0 {
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
