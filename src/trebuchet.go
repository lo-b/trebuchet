package trebuchet

import (
	"strconv"
	"strings"
	"unicode"
)

// decipherFirstDigit tries to find the first digit within a line from the
// start. If it is found the digit's rune, and it's position are returned. When
// no digit is found the position -1 and first digit 0 are returned.
func decipherFirstDigit(line string) (int, rune) {
	for position, character := range line {
		if unicode.IsDigit(character) {
			return position, character
		}
	}

	return -1, 0
}

// decipherSecondDigit tries to find a digit within a line starting from the
// end. When one is found, the rune of the digit is returned.
func decipherSecondDigit(line string) rune {
	for i := len(line) - 1; i >= 0; i-- {
		character := rune(line[i])

		if unicode.IsDigit(character) {
			return character
		}
	}

	return 0
}

func DecipherCalibrationDocument(calibrationDocument string) int {
	lines := strings.Split(calibrationDocument, "\n")
	calibrationValuesSum := 0

	for i := 0; i < len(lines); i++ {
		newStart, firstRuneDigit := decipherFirstDigit(lines[i])
		secondRuneDigit := decipherSecondDigit(lines[i][newStart:])

		decipheredStringValue := string(firstRuneDigit) + string(secondRuneDigit)

		if decipheredValue, err := strconv.Atoi(decipheredStringValue); err == nil {
			calibrationValuesSum += decipheredValue
		}
	}

	return calibrationValuesSum
}
