package main

import (
	"github.com/encima/openape"
)

func main() {
	o := openape.NewServer("$HOME/dev/go/c3s-oape/config")
	o.RunServer()
}
