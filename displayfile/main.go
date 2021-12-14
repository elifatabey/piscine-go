package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	lenght := 0
	filename := ""
	for i, prog := range os.Args {
		lenght++
		if i == 1 {
			filename = prog
		}
	}
	if lenght < 2 {
		fmt.Println("File name missing")
		return
	}
	if lenght > 2 {
		fmt.Println("Too many arguments")
		return
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(string(data))
}
