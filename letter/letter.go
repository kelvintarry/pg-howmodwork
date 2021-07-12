// This source file is in the package (directory) "letter". Note, the
// directory name and the package name are the same. Although they can
// be different names, it is better to use the same name for REUSABLE
// packages (best practice).
//
// The package name is NOT "main", therefore, it is a REUSABLE package.
//
// If we added more source files in this directory they would all have to
// include the line "package letter".
//
// The GLOBAL variable identifier "firstLetter", which begins with a
// LOWERCASE letter, is ONLY visable WITHIN the package (source files in
// this directory).
//
// The GLOBAL variable identifier "LastLetter", which begins with an
// UPPERCASE letter, is an EXPORTED IDENTIFIER, and is therefore visable
// OUTSIDE of the package.
//
// The function identifier "DisplayFirstLetter" is also an exported
// identifier and is part of the PUBLIC INTERFACE of the PACKAGE.

package letter

import "fmt"

var (
	firstLetter byte = 'A'
	LastLetter  byte = 'Z'
)

func DisplayFirstLetter() {
	fmt.Println("First letter is", string(firstLetter))
}

func DisplayName() {
	fmt.Println("Package owner: Kelvin")
}