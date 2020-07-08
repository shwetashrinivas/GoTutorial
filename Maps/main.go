package main

import "fmt"

//Maps are used to represent a collection of related properties
// All keys and all values must be of same type.
// Keys and indexed and can be iterated over
// Maps are reference type

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	/*
		colors := make(map[string]string)

	*/

	//No need to know all the keys at compile time
	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
