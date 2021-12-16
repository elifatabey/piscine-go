package main

import (
	"fmt"
	"os"
	"strconv"
)

func signcheck(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func main() {
	sign := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]
	if len(args) != 3 {
	} else {
		if signcheck(args[1], sign) {
			nb1, err := strconv.Atoi(args[0])
			nb2, err2 := strconv.Atoi(args[2])
			if nb1 >= 9223372036854775807 || nb1 <= -9223372036854775807 || nb2 >= 9223372036854775807 || nb2 <= -9223372036854775807 {
			} else {
				if err == nil && err2 == nil {
					switch args[1] {
					case "+":
						result := nb1 + nb2
						if result-nb2 != nb1 {
						} else {
							fmt.Println(nb1 + nb2)
						}
					case "-":
						result := nb1 - nb2
						if result+nb2 != nb1 {
						} else {
							fmt.Println(nb1 - nb2)
						}
					case "/":
						if nb2 == 0 {
							fmt.Println("No division by 0")
						} else {
							result := nb1 / nb2
							if result*nb2 != nb1 {
							} else {
								fmt.Println(result)
							}
						}
					case "%":
						if nb2 == 0 {
							fmt.Println("No Modulo by 0")
						} else {
							fmt.Println(nb1 % nb2)
						}
					case "*":
						result := nb1 * nb2
						if result/nb2 != nb1 {
						} else {
							fmt.Println(nb1 * nb2)
						}
					}
				} else {
					fmt.Println("1")
				}
			}
		} else {
			fmt.Println("0")
		}
	}
}
