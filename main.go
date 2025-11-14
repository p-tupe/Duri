package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
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

	logFileURL := storage.NewFileURI("./fuellogs.csv")

	logFile, err := storage.Appender(logFileURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()

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
			dump := fmt.Appendf(make([]byte, 0), "%s,%s,%s\n", odoInput.Text, fuelInput.Text, dateInput.Text)
			_, err = logFile.Write(dump)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	w.Resize(fyne.NewSize(512, 512))
	w.SetContent(form)
	w.ShowAndRun()
}
