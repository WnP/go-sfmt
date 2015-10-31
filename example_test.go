package sfmt_test

import (
	"fmt"

	"github.com/WnP/go-sfmt"
)

func ExampleCamelCased() {
	fmt.Println(sfmt.CamelCased("The_quick_brown_fox_jumps_over_the_lazy_dog"))
	// Output: theQuickBrownFoxJumpsOverTheLazyDog
}

func ExampleDashed() {
	fmt.Println(sfmt.Dashed("The quick brown fox jumps over the lazy dog"))
	// Output: the-quick-brown-fox-jumps-over-the-lazy-dog
}

func ExampleDoted() {
	fmt.Println(sfmt.Doted("The-quick-brown-fox-jumps-over-the-lazy-dog"))
	// Output: the.quick.brown.fox.jumps.over.the.lazy.dog
}

func ExampleSpaced() {
	fmt.Println(sfmt.Spaced("The.quick.brown.fox.jumps.over.the.lazy.dog"))
	// Output: the quick brown fox jumps over the lazy dog
}

func ExampleUnderscored() {
	fmt.Println(sfmt.Underscored("TheQuickBrownFoxJumpsOverTheLazyDog"))
	// Output: the_quick_brown_fox_jumps_over_the_lazy_dog
}
