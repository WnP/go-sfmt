package sfmt

import (
	"reflect"
	"testing"
)

func testEqual(t *testing.T, output string, expected string) {
	if output != expected {
		t.Errorf("found: %s\n       expected: %s", output, expected)
	}
}

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

func TestSpaced(t *testing.T) {
	testEqual(t, Spaced("from.doted.format"), "from doted format")
}

func TestDoted(t *testing.T) {
	testEqual(t, Doted("from spaced format"), "from.spaced.format")
}

func TestDashed(t *testing.T) {
	testEqual(t, Dashed("from spaced format"), "from-spaced-format")
}

func TestUnderscored(t *testing.T) {
	testEqual(t, Underscored("from spaced format"), "from_spaced_format")
}

func TestCamelCased(t *testing.T) {
	testEqual(t, CamelCased("from spaced format"), "fromSpacedFormat")
}
