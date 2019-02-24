package mod2

import (
	"fmt"

	"github.com/johndelavega/gomod"
)

var _version = gomod.Version()
var _internalVersion = "v0.0.0"

// Gomod2 used for testing replace and local modules/packages/libraries
func Gomod2() string {
	fmt.Println(fmt.Sprintf("4 mod2/Gomod2 %s - main.go import anymod 'github.com/johndelavega/gomod'", _version))
	return fmt.Sprintf("5 mod2/Gomod2 %s - main.go import anymod 'github.com/johndelavega/gomod'", _version)
}

// Gomod test
func Gomod() string {
	return fmt.Sprintf("mod2/Gomod() Version %s, internalVersion %s", _version, _internalVersion)

}
