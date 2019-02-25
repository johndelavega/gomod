//
// https://github.com/golang/go/wiki/Modules
//
// A module is a collection of related Go packages that are versioned together as a single unit.
//

package gomod

import "fmt"

const _moduleVersion = "v0.0.14" // primary project version

// todo:update name to packageVersion?
const _internalVersion = "v0.0.1" // update only when changes occurs in this package, except _moduleVersion

// Version exported
func Version() string {
	return _moduleVersion
}

// Gomod used for testing go.mod replace and local modules/packages
func Gomod() string {
	fmt.Println(fmt.Sprintf("1 Gomod() %s - main.go import anymod 'github.com/johndelavega/gomod'", _moduleVersion))
	fmt.Println(fmt.Sprintf("2 go.mod - replace github.com/johndelavega/gomod %s => ./anymod4", _moduleVersion))
	return fmt.Sprintf("gomod/Gomod() %s - return string value", _moduleVersion)
}

// FuncTest test
func FuncTest() string {
	return fmt.Sprintf("gomod/FuncTest() Module Version %s, internalVersion %s ", _moduleVersion, _internalVersion)
}
