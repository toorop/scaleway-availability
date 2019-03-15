package main

import (
	"fmt"
	"os"

	"github.com/toorop/scaleway-availability"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: scwa REGION FLAVOR\nEx: scwa par1 START1-L\n")
		os.Exit(1)
	}
	region := os.Args[1]
	flavor := os.Args[2]
	a, err := scwa.GetAvailability(region, flavor)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(a)

}
