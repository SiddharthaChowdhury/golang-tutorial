package main

import "fmt"

func main () {

	/*
		colors := map[string]string{
			"white": "#fff",
			"black": "#000",
		}
	*/

	//  Creating map
	numbers := make(map[int]string)
	colors := make(map[string]string)

	// Adding
	numbers[0] = "Zero"
	fmt.Println(numbers)
	fmt.Println(numbers[0])

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] 	= "#ff0000"
	colors["green"]	= "#4bf745"
	fmt.Println(colors)
	fmt.Println(colors["white"])
	// NOTE: `colors.white` will not work with map

	// Deleting
	delete(colors, "white")
	fmt.Println(colors)

	// Iterate
	for key, value := range colors {
		fmt.Println( key + " - " + value )
	}
}