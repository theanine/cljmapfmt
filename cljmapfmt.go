package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printNewLine(scope int) {
	fmt.Printf("\n")
	for i := 0; i < scope; i++ {
		fmt.Printf("  ")
	}
}

func main() {
	fileName := os.Args[1]

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Couldn't read file:", err)
	}

	scope := 0
	for i := 0; i < len(dat); i++ {
		if dat[i] == '{' {
			fmt.Printf("{")
			scope++
			continue
		}
		if dat[i] == '}' {
			scope--
			printNewLine(scope)
			fmt.Printf("}")
			continue
		}
		if dat[i] == ':' {
			printNewLine(scope)
			fmt.Printf(":")
			continue
		}
		fmt.Printf("%c", dat[i])
	}
	fmt.Println()
}
