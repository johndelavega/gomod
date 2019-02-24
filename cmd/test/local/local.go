package local

import (
	"fmt"

	"github.com/johndelavega/gomod"
)

var _version = gomod.Version()
var _localVersion = "v0.0.1"

// GomodLocal used for testing replace and local modules/packages/libraries
func GomodLocal() string {
	fmt.Println(fmt.Sprintf("gomod/cmd/test/local/GomodLocal() Version: %s, localVersion: %s", _version, _localVersion))
	return fmt.Sprintf("gomod/cmd/test/local/GomodLocal() Version: %s, localVersion: %s", _version, _localVersion)
}

// Gomod test
func Gomod() string {
	return fmt.Sprintf("gomod/cmd/test/local/Gomod() Version %s, _localVersion %s", _version, _localVersion)

}

// FuncTest test
func FuncTest() string {
	return fmt.Sprintf("gomod/cmd/test/local/FuncTest() Module Version %s, _localVersion %s ", _version, _localVersion)
}
