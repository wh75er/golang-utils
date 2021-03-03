package main

import (
	"calc/calculator"
	"calc/iop"
	"log"
	"os"
)

func main() {
	s, e := iop.ReadInput(os.Stdin)

	if e != nil {
		log.Fatal(e)
	}

	calculator.Calculate(s)

	s = "You string is: " + s

	iop.WriteInput(s, os.Stdout)
}
