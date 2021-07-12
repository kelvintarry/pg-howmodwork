// This source file is part of an EXECUTABLE package. All the GLOBAL
// identifiers are PRIVATE, therefore, they can be accessed from WITHIN
// source files in the same package, but NOT from within source files
// OUTSIDE of the package.

package main

import "fmt"

const (
	byteMin byte = 0 // alias for uint8
	byteMax byte = 255
)

func displayByteMin() {
	fmt.Printf("Byte type is %T\n", byteMin)
	fmt.Println("Byte minimum value is", byteMin)
}
