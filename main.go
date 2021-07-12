// This source file is in the package "main". Therefore, it is an EXECUTABLE
// package. Note, the directory name and the package name are NOT the same
// for executable packages.
//
// This source contains a "main" function, therefore, it is the beginning
// point of the program.

package main

import (
	"fmt"
	"github.com/kelvintarry/pg-howmodwork/letter"
)

func main() {
	letter.DisplayName()

	// ALL the GLOBAL identifiers beginning with an uppercase letter or
	// lowercase letter are visible WITHIN the package.
	displayByteMin()
	fmt.Println("Byte maximum value is", byteMax)

	// Only the GLOBAL identifiers beginning with an uppercase letter are
	// visible OUTSIDE of the "letter" package (Public Interface).
	letter.DisplayFirstLetter()
	fmt.Println("Last letter is", string(letter.LastLetter))
}
