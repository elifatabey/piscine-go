package piscine

func AppendRange(min, max int) []int {

	var result []int
	l := (max - min)

	if min > max {
		return nil
	} else {
		for i := 0; i < l; i++ {
			result = append(result, min+i)
		}
	}
	return result
}
