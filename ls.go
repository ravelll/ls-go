package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		dir, _ := os.Getwd()
		files, err := ioutil.ReadDir(dir)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		for _, file := range files {
			fmt.Printf("%s\t", file.Name())
		}
	} else {
		for _, arg := range os.Args[1:] {
			files, err := ioutil.ReadDir(arg)

			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}

			for _, file := range files {
				fmt.Printf("%s\t", file.Name())
			}
		}
	}
}
