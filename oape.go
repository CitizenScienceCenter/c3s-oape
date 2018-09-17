package main

import (
	"github.com/encima/openape"
)

func main() {
<<<<<<< HEAD
	o := openape.NewServer("/home/encima/dev/go/c3s-oape/config")
=======
	o := openape.NewServer("$HOME/dev/go/c3s-oape/config")
>>>>>>> ae91e297b7fd9f6aca703d46144be28f42dd1afb
	o.RunServer()
}
