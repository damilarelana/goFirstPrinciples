package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// define and initialize the command line variable that can be parsed
	// this allows us to use "-externalX=..." as a flag when running the code
	// i.e. go run simpleCommandLineParsing.go -externalX=1224
	internalX := flag.Int("externalX", 100, "the maximum random value")

	// tell Go to Parse the flag variable before usage
	flag.Parse()

	// Use the internal version i.e. internalX
	fmt.Println("generated random number:", rand.Intn(*internalX))
}
