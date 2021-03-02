package main

import (
	"calc/iop"
	"fmt"
	"log"
	"os"
)

func main() {
	s, e := iop.ReadInput(os.Stdin)

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("Your string is: ", s)

	s = "You string is: " + s

	iop.WriteInput(s, os.Stdout)
}
