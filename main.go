package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	var nav *container.Navigation

	a := app.New()
	w := a.NewWindow("Duri")

	logFileWriter := GetLogFileWriter()
	defer logFileWriter.Close()
	form := GetForm(logFileWriter, func() {
		nav.OnBack()
	})

	logFileReader := GetLogFileReader()
	defer logFileReader.Close()
	list := GetList(logFileReader, func() {
		nav.Push(form)
	})

	nav = container.NewNavigation(list)

	w.Resize(fyne.NewSize(512, 512))
	w.SetContent(nav)
	w.ShowAndRun()
}
