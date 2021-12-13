package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	} else {
		len := (max - min)
		result := make([]int, len)
		for i := 0; i < len; i++ {
			result[i] = min + i
		}
		return result
	}
}
