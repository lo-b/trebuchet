package trebuchet

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadCalibrationDocument(filePath string) string {
	dat, err := os.ReadFile(filePath)
	check(err)

	return string(dat)
}
