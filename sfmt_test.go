package sfmt

import (
	"reflect"
	"testing"
)

func testSliceEqual(t *testing.T, output []string, expected []string) {
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("found: %s\n       expected: %s", output, expected)
	}
}

func TestFormatRecognition(t *testing.T) {
	testSliceEqual(t, getElements("from spaced format"), []string{"from", "spaced", "format"})
	testSliceEqual(t, getElements("from.doted.format"), []string{"from", "doted", "format"})
	testSliceEqual(t, getElements("from-dashed-format"), []string{"from", "dashed", "format"})
	testSliceEqual(t, getElements("from_underscored_format"), []string{"from", "underscored", "format"})
	testSliceEqual(t, getElements("fromUnderscoredFormat"), []string{"from", "underscored", "format"})
	testSliceEqual(t, getElements("FromUnderscoredFormat"), []string{"from", "underscored", "format"})
}

func TestLowerSlice(t *testing.T) {
	testSliceEqual(
		t,
		lowerSlice([]string{"Some", "Uppercase", "Items"}),
		[]string{"some", "uppercase", "items"})
}
