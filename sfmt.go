package sfmt

import (
	"strings"
	"unicode"
)

// Get all words elements from given string and return them as a slice of strings
func getElements(input string) []string {
	switch {
	case strings.Index(input, " ") > -1: // spaced
		return strings.Split(input, " ")
	case strings.Index(input, ".") > -1: // doted
		return strings.Split(input, ".")
	case strings.Index(input, "-") > -1: // dashed
		return strings.Split(input, "-")
	case strings.Index(input, "_") > -1: // underscored
		return strings.Split(input, "_")
	default: // CamelCase
		output := []string{}
		i := 0
		for s := input; s != ""; s = s[i:] {
			i = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
			if i <= 0 {
				i = len(s)
			}
			output = append(output, strings.ToLower(s[:i]))
		}
		return output
	}
}

// Convert given string slice to lower string slice
func lowerSlice(input []string) []string {
	var output []string
	for _, v := range input {
		output = append(output, strings.ToLower(v))
	}
	return output
}

// Convert any string format to spaced format
// all ouput letters are lowercase
func Spaced(input string) string {
	return strings.Join(lowerSlice(getElements(input)), " ")
}

// Convert any string format to dashed format
// all ouput letters are lowercase
func Dashed(input string) string {
	return strings.Join(lowerSlice(getElements(input)), "-")
}

// Convert any string format to doted format
// all ouput letters are lowercase
func Doted(input string) string {
	return strings.Join(lowerSlice(getElements(input)), ".")
}

// Convert any string format to underscored format
// all ouput letters are lowercase
func Underscored(input string) string {
	return strings.Join(lowerSlice(getElements(input)), "_")
}

// Convert any string format to underscored format
// first letter is lowercase
func CamelCased(input string) string {
	var output string
	e := lowerSlice(getElements(input))
	for k, v := range e {
		if k == 0 {
			output += v
		} else {
			output += strings.Title(v)
		}
	}
	return output
}
