package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today+2:
		fmt.Println("Aftertomorrow")
	case today+1:
		fmt.Println("Tomorrow")
	case today+0:
		fmt.Println("Today")
	default:
		fmt.Println("far away")
	}

	switch  {
	case time.Now().Hour() < 12:
		fmt.Println("Now is noon")
	case time.Now().Hour() < 18:
		fmt.Println("Now is afternoon")
	default:
		fmt.Println("Now is night")
	}
}
