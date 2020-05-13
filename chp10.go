package main

import (
	"fmt"
	"log"
)

type Location struct {
	Longitude float64
	Latitude  float64
	Name      string
}

func (d *Location) SetName(name string) error {

	d.Name = name
	return nil
}

func (d *Location) SetLatitude(lat float64) error {

	d.Latitude = lat
	return nil
}

func (d *Location) SetLongitude(lon float64) error {

	d.Longitude = lon
	return nil
}
func main() {

	location := Location{}
	err := location.SetName("Mumbai")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(25.98)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(96.96)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name, location.Longitude, location.Latitude)

}
