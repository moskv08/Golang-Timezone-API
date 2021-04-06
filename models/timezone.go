package models

import (
	"time"
)

type Timezone struct {
	Time string
	UTC  string
	Zone string
}

// Method
func (tz *Timezone) FindProperZone(zone string) Timezone {

	location, err := time.LoadLocation(zone)
	if err != nil {
		panic(err)
	}

	tz.UTC = time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC).String()
	tz.Time = time.Now().In(location).String()
	tz.Zone = location.String()

	// tz := Timezone{time.String(), timeInUTC.String(), location.String()}
	return *tz
}

// Method
func (tz *Timezone) FindMyZone() Timezone {
	time := time.Now()
	zone, _ := time.Zone()

	tz.Time = time.String()
	tz.Zone = zone

	return *tz
}
