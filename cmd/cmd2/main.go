//
// cmd/cmd2
//
package main

import (
	"fmt"

	"github.com/johndelavega/gomod"
	"github.com/johndelavega/gomod/mod2"
)

var _version = gomod.Version()

func main() {
	fmt.Println(fmt.Sprintf("cmd2(.exe) %s\n", _version))

	fmt.Println(fmt.Sprintf("gomod/cmd/cmd2/main.go %s\n", _version))
	fmt.Println(fmt.Sprintf("gomod.Gomod(): %s\n", gomod.Gomod()))
	fmt.Println(fmt.Sprintf("mod2.Gomod(): %s\n", mod2.Gomod2()))
}
