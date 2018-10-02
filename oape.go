package main

import (
	"github.com/encima/openape"
)

func main() {
	o := openape.NewServer("config")
	o.RunServer()
}
