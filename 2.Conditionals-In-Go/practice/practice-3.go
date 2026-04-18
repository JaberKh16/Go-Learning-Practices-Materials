package main

import (
	"fmt"
	"time"
)


func main() {
	weekDay := time.Now().Weekday()

	switch weekDay {
	case time.Saturday, time.Sunday:
		fmt.Println("weekday")

	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Println("office day")
		
	case time.Friday:
		fmt.Println("Friday.........")
		
	default:
		fmt.Println("invalid..")
	}
}