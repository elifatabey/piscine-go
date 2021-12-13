package piscine

func MakeRange(min, max int) []int {
	if min > max || min == 0 && max == 0 {
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
