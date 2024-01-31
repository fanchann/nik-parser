package main

import (
	"fmt"

	nikparser "github.com/fanchann/nik-parser"
)

func main() {
	// nikOpts := nikparser.NIKOpts{District: "BEKASI TIMUR", PostalCode: "17111"}
	niks := nikparser.NIKGen()
	fmt.Printf("niks: %v\n", niks)
}
