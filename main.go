package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("Pep OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	myApp.Settings().SetTheme(theme.LightTheme())

	img = canvas.NewImageFromFile("C:\\Users\\SANTOSH\\Desktop\\os\\os.jpg")

	btn1 = widget.NewButtonWithIcon("Calculator app", theme.ContentAddIcon(), func() {
		showCalc()
	})

	btn2 = widget.NewButtonWithIcon("Gallery app", theme.StorageIcon(), func() {
		showGallery()
	})

	btn3 = widget.NewButtonWithIcon("editor app", theme.DocumentIcon(), func() {
		showEditor()
	})

	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(4, DeskBtn, btn1, btn2, btn3)))

	myWindow.Resize((fyne.NewSize(1280, 720)))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()

}
