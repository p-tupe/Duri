package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Duri")

	logFileWriter := GetLogFileWriter()
	defer logFileWriter.Close()
	_ = GetForm(logFileWriter)

	logFileReader := GetLogFileReader()
	defer logFileReader.Close()
	list := GetList(logFileReader)

	w.Resize(fyne.NewSize(512, 512))
	w.SetContent(list)
	w.ShowAndRun()
}
