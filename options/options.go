package options

import (
	"fmt"
	"os/exec"
	"time"
)

func MeasureCommand(lines []string) {
	s := time.Now()

	out, err := exec.Command(lines[1], lines[2:]...).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	e := time.Now().Sub(s).Nanoseconds()

	fmt.Println("time:", e, "ns")
}