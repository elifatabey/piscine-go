package piscine

func TrimAtoi(s string) int {
	o_number := 0
	c := 0

	a_s := []rune(s)
	pl := 1
	minus_pos := 0
	minus_have := false
	first_num := 0

	for index, num := range a_s {
		if num == '-' {
			minus_pos = index
			minus_have = true
		}
	}
	for _, word := range a_s {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		}
	}
	if minus_have {
		for index, num := range a_s {
			if num >= '0' && num <= '9' {
				first_num = index
				break
			}
		}
	}
	if first_num > minus_pos {
		pl = -1
	}
	return o_number * pl
}

// func TrimAtoi(s string) int {
// 	RuneS := []rune(s)
// 	count := 0
// 	numbers := []rune("0123456789")
// 	for i := 0; i <= 9; i++ {
// 		for k := 0; k < len(RuneS); k++ {
// 			if rune(s[k]) == numbers[i] {
// 				count = count + 1
// 			}
// 		}
// 	}
// 	if count == 0 {
// 		return 0
// 	} else {
// 		number := []rune{}
// 		for i := 0; i <= 9; i++ {
// 			for k := 0; k < len(RuneS); k++ {
// 				if RuneS[k] == numbers[i] {
// 					number = append (number, numbers[i])
// 				}
// 			}
// 		}
// 		len := len(number)
// 		var jum int = 1
// 		var result int = 0
// 		for i:=len; i>0;i-- {
// 			jum = int((number[i] - rune(48)))
// 			result = result + (jum*(len-1))
// 			jum = jum*10
// 		}
// 		return result
// 	}
// }

// // else {
// // 	if RuneS[i] == '-' {
// // 		ERune = append(ERune, RuneS[i])
// // 	}
// // }
