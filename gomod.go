//
// https://github.com/golang/go/wiki/Modules
//
// A module is a collection of related Go packages that are versioned together as a single unit.
//

package gomod

import "fmt"

const _version = "v0.0.5"

// Version exported
func Version() string {
	return _version
}

// Gomod use for testing replace and local modules/packages
func Gomod() string {
	fmt.Println(fmt.Sprintf("1 Gomod %s - main.go import anymod 'github.com/johndelavega/gomod'", _version))
	fmt.Println(fmt.Sprintf("2 go.mod - replace github.com/johndelavega/gomod %s => ./anymod4", _version))
	return fmt.Sprintf("3 Gomod %s - main.go import anymod4 'github.com/johndelavega/gomod'", _version)
}
