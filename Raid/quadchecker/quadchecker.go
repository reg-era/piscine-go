package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func quadsrch(fileName, x, y, str string) bool {
	cmd := exec.Command(fileName, x, y)
	output, _ := cmd.CombinedOutput()

	return (str == string(output))
}

func Quadcheck(str string) {
	allquad := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}

	y, x := 0, 0
	endX := false

	for _, char := range str {
		if char == '\n' {
			endX = true
			y++
		}
		if !endX {
			x++
		}
	}

	xa := strconv.Itoa(x)
	yb := strconv.Itoa(y)

	var res []map[string]string

	for _, quad := range allquad {

		if quadsrch(quad, xa, yb, str) {
			quadfind := make(map[string]string)

			quadfind["quadnm"] = quad
			quadfind["x"] = xa
			quadfind["y"] = yb
			res = append(res, quadfind)
		}

	}

	if len(res) < 1 || str == "" {
		fmt.Println("Not a quad function")
		return
	}

	for i, item := range res {
		fmt.Printf("[%v] [%v] [%v]", item["quadnm"][2:], item["x"], item["y"])
		if i < len(res)-1 {
			fmt.Print(" || ")
		}
	}

	fmt.Println()
}

func main() {
	arg := os.Args[1:]
	Quadcheck(arg[0])
}
