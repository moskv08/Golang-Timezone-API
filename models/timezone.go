package models

import (
	"fmt"
	"time"
)

type Timezone struct {
	Zone string
	Time string
}

func GetSpecificTimeZone(location string) time.Time {
	// Check for three char
	if len(location) != 3 {
		fmt.Println("Yo! Just three characters G")
	}

	loc, _ := time.LoadLocation(location)
	time := time.Now().In(loc)

	return time
}

func GetMyTimeZone() (string, time.Time) {
	time := time.Now()
	zone, _ := time.Zone()
	// fmt.Println("ZONE : ", z, " Time : ", time) // local time

	return zone, time
}
