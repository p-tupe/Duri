package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var logFileURL = storage.NewFileURI("./fuellogs.csv")

func GetLogFileWriter() fyne.URIWriteCloser {
	writer, err := storage.Appender(logFileURL)
	if err != nil {
		log.Fatalln(err)
	}

	return writer
}

func GetLogFileReader() fyne.URIReadCloser {
	reader, err := storage.Reader(logFileURL)
	if err != nil {
		log.Fatalln(err)
	}

	return reader
}
