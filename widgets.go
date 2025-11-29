package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetForm(logFile fyne.URIWriteCloser, showList func()) *widget.Form {
	odoInput := widget.NewEntry()
	fuelInput := widget.NewEntry()
	dateInput := &widget.DateEntry{}

	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Odometer (km)", Widget: odoInput},
			{Text: "Fuel (ltr)", Widget: fuelInput},
			{Text: "Date", Widget: dateInput},
		},
		OnSubmit: func() {
			dump := fmt.Appendf(make([]byte, 0), "%s,%s,%s\n", odoInput.Text, fuelInput.Text, dateInput.Text)
			_, err := logFile.Write(dump)
			if err != nil {
				log.Fatalln(err)
			}
			showList()
		},
		OnCancel: func() {
			showList()
		},
	}
}

func GetList(logFile fyne.URIReadCloser, showForm func()) *fyne.Container {
	fls := []FuelLog{}

	reader := csv.NewReader(logFile)
	for {
		row, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatalln(err)
			}
		}
		fls = append(fls, NewFuelLog(row[0], row[1], row[2]))
	}

	list := widget.NewList(
		func() int {
			return len(fls)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("All Logs")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(fls[i].String())
		})

	btn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		showForm()
	})

	return container.NewVBox(list, btn)
}
