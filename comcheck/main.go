package main

import (
	"fmt"
	"os"
)

func main() {
	alertif := []string{"01", "galaxy", "galaxy 01"}
	for i := 0; i <= 2; i++ {
		for _, str := range os.Args {
			if str == alertif[i] {
				fmt.Println("Alert!!!")
			}
		}
	}
}
