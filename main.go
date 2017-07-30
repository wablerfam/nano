package main

import (
	"fmt"
	"os"

	"nano/nano"
	"nano/options"
)

func main() {
	if len(os.Args) != 1 {
		options.MeasureCommand(os.Args)
		return
	}

	n := nano.Measure()

	fmt.Println(n)
}
