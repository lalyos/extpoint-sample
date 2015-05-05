package main

import (
	"fmt"

	"github.com/lalyos/extpoint-sample/extpoints"
)

func init() {
	extpoints.Register(new(HelloComponent), "hello")
}

type HelloComponent struct {
}

func (c *HelloComponent) Run(args []string) {
	fmt.Println("Hello world..")
}
