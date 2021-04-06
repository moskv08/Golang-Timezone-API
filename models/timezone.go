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

func (tz *Timezone) FindMyZone() (string, string) {
	time := time.Now()
	zone, _ := time.Zone()
	// fmt.Println("ZONE : ", z, " Time : ", time) // local time
	return time.String(), zone
}
