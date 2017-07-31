package main

import (
	"fmt"
	"os"
	"runtime"

	"nano/nano"
	"nano/options"
)

func main() {
	if len(os.Args) == 1 {
		n := nano.Measure()
		fmt.Println(n)

	} else if len(os.Args) != 1 {
		if runtime.GOOS == "linux" {
			options.MeasureCommand(os.Args)	
					
		} else {
			fmt.Println("Error: command line argument option is valid only for linux")
		}
	}
}
