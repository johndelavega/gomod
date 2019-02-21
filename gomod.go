package gomod

import "fmt"

// Gomod use for testing replace and local modules
func Gomod() string {
	fmt.Println("Gomod() - import 'github.com/johndelavega/gomod'")
	return "Gomod() - import 'github.com/johndelavega/gomod v0.0.1'"
}
