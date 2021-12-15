package main

import (
	"fmt"
	"os"
)

func main() {
	alertif := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	for i := 0; i <= 2; i++ {
		for _, str := range os.Args {
			if str == alertif[i] {
				count++
			}
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
