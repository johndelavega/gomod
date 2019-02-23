package mod2

import (
	"fmt"

	"github.com/johndelavega/gomod"
)

var _version = gomod.Version()

// Gomod used for testing replace and local modules/packages/libraries
func Gomod() string {
	fmt.Println(fmt.Sprintf("Gomod %s - main.go import anymod 'github.com/johndelavega/gomod'", _version))
	fmt.Println(fmt.Sprintf("go.mod - replace github.com/johndelavega/gomod %s => ./anymod", _version))
	return fmt.Sprintf("Gomod %s - main.go import anymod 'github.com/johndelavega/gomod'", _version)
}
