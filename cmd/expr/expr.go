package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gonutz/expr"
)

func main() {
	calculator := expr.NewCalculator()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		x, err := calculator.Evaluate(line)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(x)
		}
	}
}
