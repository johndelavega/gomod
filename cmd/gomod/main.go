//
// cmd/gomod
//
package main

import (
	"fmt"

	"github.com/johndelavega/gomod"
	"github.com/johndelavega/gomod/mod2"
)

var _version = gomod.Version()

func main() {
	fmt.Println(fmt.Sprintf("gomod(.exe) %s\n", _version))

	fmt.Println(fmt.Sprintf("gomod/cmd/gomod/main.go %s\n", _version))
	fmt.Println(fmt.Sprintf("gomod.Gomod(): %s\n", gomod.Gomod()))
	fmt.Println(fmt.Sprintf("mod2.Gomod(): %s\n", mod2.Gomod2()))
}
