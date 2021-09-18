package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"I", "am", "stupid", "and", "weak"}

	fmt.Printf("mySlice old %+v\n", mySlice)

	for index := range mySlice {
		if index == 2 {
			mySlice[index] = "smart"
		} else if index == 4 {
			mySlice[index] = "strong"
		}
	}

	fmt.Printf("mySlice new %+v\n", mySlice)
}
