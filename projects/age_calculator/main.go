package main

import (
	"fmt"
	"time"
)

func main() {
	birthdate := time.Date(1990, 6, 1, 12, 0, 0, 0, time.UTC)
	fmt.Printf(
		"The age of a person born on %s is %d.\n",
		birthdate.Format("2006-01-02"), CalculatorAge(birthdate),
	)
}

func CalculatorAge(birthdate time.Time) int {
	return int(time.Now().Sub(birthdate).Hours() / 24 / 365)
}
