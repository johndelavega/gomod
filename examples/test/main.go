//
// cmd/test
//
package main

import (
	"fmt"

	// _ "./local"
	"github.com/johndelavega/gomod"
	// "github.com/johndelavega/gomod/mod2"
	using__replace__mod2_to_local "github.com/johndelavega/gomod/mod2"
)

const _appVersion = "v0.0.1"

var _version = gomod.Version()

func main() {
	fmt.Println(fmt.Sprintf("gomod/cmd/test(.exe) %s (Primary Package / Module Version %s)\n", _appVersion, _version))

	fmt.Println(fmt.Sprintf("gomod/cmd/test/main.go %s\n", _appVersion))
	fmt.Println(fmt.Sprintf("gomod.Gomod(): %s\n", gomod.Gomod()))
	// fmt.Println(fmt.Sprintf("mod2.Gomod(): %s\n", mod2.Gomod2()))
	// fmt.Println(fmt.Sprintf("local.GomodLocal(): %s\n", using__replace__mod2_to_local.GomodLocal()))
	fmt.Println(fmt.Sprintf("local.Gomod(): %s\n", using__replace__mod2_to_local.Gomod()))
	fmt.Println(fmt.Sprintf("local.FuncTest(): %s\n", using__replace__mod2_to_local.FuncTest()))
}
