package models

import (
	"io/ioutil"
	"strings"
	"time"
)

type Timezone struct {
	Time   string
	Offset int
	UTC    string
	Zone   string
}

// Method
func (tz *Timezone) FindProperZone(zone string) Timezone {

	// location, err := time.LoadLocation(strings.ToUpper(zone))
	location, err := time.LoadLocation(zone)
	if err != nil {
		panic(err)
	}

	tz.UTC = time.Now().UTC().String()
	tz.Time = time.Now().In(location).String()
	tz.Zone, tz.Offset = time.Now().In(location).Zone()

	return *tz
}

// Method
func (tz *Timezone) FindMyZone() Timezone {

	time := time.Now()
	tz.Zone, tz.Offset = time.Zone()
	tz.UTC = time.UTC().String()
	tz.Time = time.String()

	return *tz
}

func (tz *Timezone) ListOfLocations() *[]string {

	var zoneDir string = "/usr/share/zoneinfo"
	result := ReadFile(zoneDir)

	return result
}

func ReadFile(path string) *[]string {
	var list []string
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			//fmt.Println(f.Name())
			list = append(list, f.Name())
			ReadFile(path + "/" + f.Name())
		} else {
			//fmt.Println((path + "/" + f.Name())[1:])
			list = append(list, f.Name())
		}
	}

	return &list
}
