package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/toorop/scaleway-availability"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: scwan REGION FLAVOR\nEx: scwan par1 START1-L\n")
		os.Exit(1)
	}
	region := os.Args[1]
	flavor := os.Args[2]
	for {
		a, err := scwa.GetAvailability(region, flavor)
		if err != nil {
			beeep.Alert("scawn error", err.Error(), "")
		}
		if a != "shortage" {
			notify("Scaleway Availability", fmt.Sprintf("Hurra %s is available at %s !", flavor, region))
			break
		}
		time.Sleep(5 * time.Minute)
	}
}

func notify(title, message string) error {
	if err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration); err != nil {
		return err
	}
	return beeep.Notify(title, message, "")
}
