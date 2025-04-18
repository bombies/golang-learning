package main

import (
	"ajani.me/golang-learning/arrays"
	"ajani.me/golang-learning/loops"
	"ajani.me/golang-learning/structs"
)

func main() {
	structs.TestStructs()
	structs.F()
	structs.UsingEmbeddedStructs()
	structs.UsingStructMethods()
	structs.TestNewUser()
	structs.TestSendMessage()

	// Loops
	loops.FizzBuzz()

	// Arrays
	arrays.TestCurryingChallenge()
}
