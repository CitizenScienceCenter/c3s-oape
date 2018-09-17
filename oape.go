package main

import (
	"github.com/encima/openape"
)

func main() {
	o := openape.NewServer("/home/encima/dev/go/c3s-oape/config")
	o.RunServer()
}
