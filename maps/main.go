package main

import "fmt"

func main() {
	empty_map := make(map[string]string)
	fmt.Println(empty_map)

	colors := map[string]string{
		"red":   "#ff0000",
	}
	fmt.Println(colors)

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	for key, value := range colors {
		fmt.Println(key, value)
	}
}
