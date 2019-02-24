package local

import (
	"fmt"

	"github.com/johndelavega/gomod"
)

var _version = gomod.Version()
var _localVersion = "v0.0.0"

// GomodLocal used for testing replace and local modules/packages/libraries
func GomodLocal() string {
	fmt.Println(fmt.Sprintf("local/GomodLocal() Version: %s, localVersion: %s", _version, _localVersion))
	return fmt.Sprintf("local/GomodLocal() Version: %s, localVersion: %s", _version, _localVersion)
}
