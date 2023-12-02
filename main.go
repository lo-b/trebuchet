package main

import (
	"fmt"
	"github.com/lo-b/trebuchet/src"
	"strings"
)

func main() {
	calibrationDocument := trebuchet.ReadCalibrationDocument("puzzle-input.txt")
	calibrationValue := trebuchet.DecipherCalibrationDocument(calibrationDocument)

	println(strings.Repeat("-", 80))
	println("\tTREBUCHET OUTPUT")
	println(strings.Repeat("-", 80))
	fmt.Printf("\tCalibration sum:\t%d", calibrationValue)
}
