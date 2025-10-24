package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	w := app.New().NewWindow("Duri")

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
		},
	}

	w.SetContent(form)
	w.ShowAndRun()
}
