package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	for _, word := range args {
		for _, r := range word {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				return
			}
		}
	}
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[2])
	if a > 1000 {
		return
	}
	if b > 1000 {
		return
	}
	// fmt.Print((args[0]))
	for _, value := range args {
		if value == "+" {
			result := a + b
			Itoa(result)
		}
		if value == "-" {
			result := a - b
			Itoa(result)
		}
		if value == "*" {
			result := a * b
			Itoa(result)
		}
		if value == "/" {
			if b == 0 {
				os.Stderr.WriteString("No division by 0")
				os.Stderr.WriteString("\n")
				return
			}
			result := a / b
			Itoa(result)
		}
		if value == "%" {
			if b == 0 {
				os.Stderr.WriteString("No modulo by 0")
				os.Stderr.WriteString("\n")
				return
			}
			result := a % b
			Itoa(result)
		}
	}
}

func Itoa(n int) {
	output := []rune{}
	sign := 1
	if n < 0 {
		sign = -1
	}
	if n == 0 {
		os.Stderr.WriteString("0")
		os.Stderr.WriteString("\n")
	}
	if n != 0 {
		n = n * (sign)
		for n != 0 {
			digit := n % 10
			output = append([]rune{rune(digit + 48)}, output...)
			n /= 10
		}
		if sign == -1 {
			output = append([]rune{'-'}, output...)
		}
		os.Stderr.WriteString(string(output))
		os.Stderr.WriteString("\n")
	}
}

// package main

// import (
// 	"os"
// 	"strconv"
// )

// func signcheck(str string, arr []string) bool {
// 	for _, v := range arr {
// 		if str == v {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	sign := []string{"+", "*", "-", "/", "%"}
// 	args := os.Args[1:]
// 	if len(args) != 3 {
// 	} else {
// 		if signcheck(args[1], sign) {
// 			nb1, err := strconv.Atoi(args[0])
// 			nb2, err2 := strconv.Atoi(args[2])
// 			// if nb1 >= 9223372036854775807 || nb1 <= -9223372036854775807 || nb2 >= 9223372036854775807 || nb2 <= -9223372036854775807 {
// 			// } else {
// 			if err == nil || err2 == nil {
// 				switch args[1] {
// 				case "+":
// 					var result int
// 					result = nb1 + nb2
// 					// if result - nb2 != nb1 {
// 					// } else {
// 					changetoString(result)
// 					// }
// 				case "-":
// 					result := nb1 - nb2
// 					if result+nb2 != nb1 {
// 					} else {
// 						changetoString(result)
// 					}
// 				case "/":
// 					if nb2 == 0 {
// 						os.Stderr.WriteString("No division by 0")
// 						os.Stderr.WriteString("\n")
// 						return
// 					} else {
// 						result := nb1 / nb2
// 						if result*nb2 != nb1 {
// 						} else {
// 							changetoString(result)
// 						}
// 					}
// 				case "%":
// 					if nb2 == 0 {
// 						os.Stderr.WriteString("No modulo by 0")
// 						os.Stderr.WriteString("\n")
// 						return
// 					} else {
// 						result := nb1 % nb2
// 						changetoString(result)
// 					}
// 				case "*":
// 					result := nb1 * nb2
// 					if result/nb2 != nb1 {
// 					} else {
// 						changetoString(result)
// 					}
// 				}
// 			} else {
// 				os.Stderr.WriteString("1")
// 				os.Stderr.WriteString("\n")
// 				return
// 			}
// 			// }
// 		} else {
// 			os.Stderr.WriteString("0")
// 			os.Stderr.WriteString("\n")
// 			return
// 		}
// 	}
// }

// func changetoString(nbr int) string {
// 	var output string
// 	output = strconv.Itoa(nbr)
// 	return output + "\n"
// }
