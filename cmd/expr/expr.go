package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gonutz/expr"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "q" {
			break
		}
		x, err := expr.Calculate(line)
		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println(x)
		}
	}
}
