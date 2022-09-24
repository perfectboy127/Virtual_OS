package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize((fyne.NewSize(400, 300)))
	output := ""
	input := widget.NewLabel(output)
	isHistory := false
	historyStr := ""
	history := widget.NewLabel(historyStr)
	var historyArr []string

	historyBtn := widget.NewButton("History", func() {
		if isHistory {
			historyStr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr = historyStr + historyArr[i]
				historyStr += "\n"
			}
		}
		isHistory = !isHistory
		history.SetText(historyStr)

	})

	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})

	clearBtm := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})

	openBtm := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	closerBtm := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})

	divideBtm := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	sevenBtm := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})

	eightBtm := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})

	nineBtm := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})

	multipluyBtm := widget.NewButton("x", func() {
		output = output + "x"
		input.SetText(output)
	})

	fourBtm := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	fiveBtm := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})

	sixBtm := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})

	minusBtm := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})

	oneBtm := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})

	twoBtm := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})

	threeBtm := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})

	plusBtm := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})

	zeroBtm := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	dotBtm := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})

	equalBtm := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}

		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewAdaptiveGrid(4,
				clearBtm,
				openBtm,
				closerBtm,
				divideBtm,
			),
			container.NewGridWithColumns(4,
				nineBtm,
				eightBtm,
				sevenBtm,
				multipluyBtm,
			),
			container.NewGridWithColumns(4,
				fourBtm,
				fiveBtm,
				sixBtm,
				minusBtm,
			),
			container.NewGridWithColumns(4,
				oneBtm,
				twoBtm,
				threeBtm,
				plusBtm,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtm,
					dotBtm,
				),
				equalBtm,
			),
		),
	))

	w.ShowAndRun()
}
