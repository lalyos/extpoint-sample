//go:generate go-extpoints
package main

import (
	"fmt"
	"os"

	"github.com/lalyos/extpoint-sample/extpoints"
)

var subcommands = extpoints.Subcommands

func usage() {
	fmt.Println("Available commands:\n")
	for name, _ := range subcommands.All() {
		fmt.Println(" - ", name)
	}
	os.Exit(2)
}

func main() {

}
