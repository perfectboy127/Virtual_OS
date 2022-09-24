package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showEditor() {

	a := app.New()
	w := a.NewWindow("pep editor")
	w.Resize(fyne.NewSize(600, 600))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("pep text editor"),
		),
	)

	content.Add(widget.NewButton("add new file", func() {
		content.Add((widget.NewLabel("new file" + strconv.Itoa(count))))
		count++
	}))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Text...")

	input.Resize(fyne.NewSize(400, 400))

	w.SetContent(
		container.NewVBox(
			content,
			input,
		),
	)

	w.ShowAndRun()
}
