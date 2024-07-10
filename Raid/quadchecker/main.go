package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input from stdin:", err)
		os.Exit(1)
	}

	cmd := exec.Command("./quadchecker", string(input))

	output, _ := cmd.CombinedOutput()

	fmt.Print(string(output))
}
