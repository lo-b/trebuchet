package trebuchet

import (
	"testing"
)

func TestDecipherCalibrationDocumentExampleInput(t *testing.T) {
	calibrationDocument := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	actual := DecipherCalibrationDocument(calibrationDocument)
	want := 142

	if actual != want {
		t.Fatalf("Failed: expected decipher calibration value %d, got %d", want, actual)
	}
}
