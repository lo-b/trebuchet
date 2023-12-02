package trebuchet

import (
	"strconv"
	"strings"
	"unicode"
)

func decipherLine(line string) int {
	firstDigit := 0
	secondDigit := 0

	for _, character := range line {
		if unicode.IsDigit(character) {
			// first digit unset
			if firstDigit == 0 {
				firstDigit = int(character)
			}
			// first digit set
			if firstDigit > 0 {
				// second digit unset
				if secondDigit == 0 {
					secondDigit = int(character)
				}

				if secondDigit > 0 {
					secondDigit = int(character)
				}
			}
		}
	}

	// if second digit is unset after loop
	// set second to fist digit
	if secondDigit == 0 {
		secondDigit = firstDigit
	}

	decipheredValue := string(rune(firstDigit)) + string(rune(secondDigit))

	if s, err := strconv.Atoi(decipheredValue); err == nil {
		return s
	}

	return 0
}

func DecipherCalibrationDocument(calibrationDocument string) int {
	lines := strings.Split(calibrationDocument, "\n")
	calibrationValue := 0
	for i := 0; i < len(lines); i++ {
		calibrationValue += decipherLine(lines[i])
	}

	return calibrationValue
}
