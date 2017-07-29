package main

import (
	"regexp"
	"testing"
)

func TestGetNanoSecondFormatRFC3339Nano(t *testing.T) {
	n := nano()

	rep := regexp.MustCompile(`^[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]T[0-9][0-9]:[0-9][0-9]:[0-9][0-9].(.*)Z`)
	mat := rep.MatchString(n)
	if mat != true {
		t.Fatalf("Not obtained nano second accurately. %v", n)
	}
}