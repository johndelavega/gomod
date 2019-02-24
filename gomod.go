//
// https://github.com/golang/go/wiki/Modules
//
// A module is a collection of related Go packages that are versioned together as a single unit.
//

package gomod

import "fmt"

const _moduleVersion = "v0.0.11" // primary project version

const _internalVersion = "v0.0.0" // update only when changes occurs in this package, except _moduleVersion

// Version exported
func Version() string {
	return _moduleVersion
}

// Gomod use for testing replace and local modules/packages
func Gomod() string {
	fmt.Println(fmt.Sprintf("1 Gomod() %s - main.go import anymod 'github.com/johndelavega/gomod'", _moduleVersion))
	fmt.Println(fmt.Sprintf("2 go.mod - replace github.com/johndelavega/gomod %s => ./anymod4", _moduleVersion))
	return fmt.Sprintf("3 Gomod() %s - main.go import anymod4 'github.com/johndelavega/gomod'", _moduleVersion)
}

// FuncTest test
func FuncTest() string {
	return fmt.Sprintf("gomod/FuncTest() Module Version %s, internalVersion %s ", _moduleVersion, _internalVersion)
}
