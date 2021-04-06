package models

import (
	"time"
)

type Timezone struct {
	Time string
	Zone string
}

// Method
func FindProperZone(zone string) Timezone {

	if len(zone) > 4 {
		//return "", "Provide proper format"
	}

	loc, _ := time.LoadLocation(zone)
	time := time.Now().In(loc)

	tz := Timezone{time.String(), loc.String()}

	return tz
}

func FindMyZone() Timezone {
	time := time.Now()
	zone, _ := time.Zone()

	tz := Timezone{time.String(), zone}

	return tz
}
