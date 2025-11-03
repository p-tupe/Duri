package main

import (
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type FuelLog struct {
	odometer int
	fuel     float64
	date     time.Time
}

func main() {
	a := app.New()
	w := a.NewWindow("Duri")

	odoInput := widget.NewEntry()
	fuelInput := widget.NewEntry()
	dateInput := &widget.DateEntry{}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Odometer (km)", Widget: odoInput},
			{Text: "Fuel (ltr)", Widget: fuelInput},
			{Text: "Date", Widget: dateInput},
		},
		OnSubmit: func() {
			log.Println("Form submitted:", odoInput.Text, fuelInput.Text, dateInput.Text)
			odometer, err := strconv.Atoi(odoInput.Text)
			if err != nil {
				panic(err)
			}
			fuel, err := strconv.ParseFloat(fuelInput.Text, 64)
			if err != nil {
				panic(err)
			}
			date, err := time.Parse("02/01/2006", dateInput.Text)
			if err != nil {
				panic(err)
			}
			fl := FuelLog{odometer, fuel, date}
			log.Println(fl)
		},
	}

	w.Resize(fyne.NewSize(512, 512))
	w.SetContent(form)
	w.ShowAndRun()
}
