package gomod

import "fmt"

// Gomod use for testing replace and local modules/packages
func Gomod() string {
	fmt.Println("Gomod v0.0.2 - main.go import anymod 'github.com/johndelavega/gomod'")
	fmt.Println("go.mod - replace github.com/johndelavega/gomod v0.0.2 => ./anymod")
	return "Gomod v0.0.2 - main.go import anymod 'github.com/johndelavega/gomod'"
}
