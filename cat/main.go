package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		cont, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			perror(err.Error())
		}
		printres(string(cont))
	} else {
		for _, file := range arg {
			cont, err := ioutil.ReadFile(file)
			if err != nil {
				perror("ERROR: " + err.Error())
			}
			printres(string(cont))
		}
	}
}

func printres(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}

func perror(err string) {
	printres(err + "\n")
	os.Exit(1)
}
