package main

import (
	"fmt"
	"strconv"
)

//! Scoped to package : initialize alot of variables
//? CAPITAL variable is GLOBAL
//*

var (
	husbandName string = "Brandon"
	husbandAge int = 43
)

func main() {
	fmt.Println("test")

	// float32 float64
	var myAge int = 37
	myAge = 38

	// simple
	simple := 20
	var convertedSimple string = strconv.Itoa(simple)

	var myName string = "Amber"
	fmt.Println("Husband info", husbandName, husbandAge)
	fmt.Println("My info", myName, myAge)
	fmt.Println("Simple format declaration", convertedSimple)
}
