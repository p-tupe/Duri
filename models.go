package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type FuelLog struct {
	odometer int
	fuel     float64
	date     time.Time
}

func NewFuelLog(odometer string, fuel string, date string) FuelLog {
	o, err := strconv.Atoi(odometer)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := strconv.ParseFloat(fuel, 64)
	if err != nil {
		log.Fatalln(err)
	}

	d, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Fatalln(err)
	}

	return FuelLog{o, f, d}
}

func (fl FuelLog) String() string {
	return fmt.Sprintf("%d,%f,%v", fl.odometer, fl.fuel, fl.date)
}
