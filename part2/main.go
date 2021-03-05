package main

import (
	"calc/calculator"
	"calc/iop"
	"log"
	"os"
	"strconv"
)

func main() {
	s, e := iop.ReadInput(os.Stdin)
	if e != nil {
		log.Fatal(e)
	}

	result, e := calculator.Calculate(s)
	if e != nil {
		log.Fatal(e)
	}

	iop.WriteInput(strconv.FormatFloat(result, 'f', 6, 64), os.Stdout)
}
