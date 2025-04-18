package main

import (
	_ "embed"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
}

//go:embed quotesData.json
var quotesData []byte

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Daily Quote")

	quoteTextLabel := widget.NewLabel("")
	quoteTextLabel.Wrapping = fyne.TextWrapWord

	quoteAuthorLabel := widget.NewLabel("")

	updateQuote := func() {
		quote := getQuote()
		quoteTextLabel.SetText("Quote: " + quote.Text)
		quoteAuthorLabel.SetText("Author: " + quote.Author)
	}

	updateQuote()

	newQuoteButton := widget.NewButton("New Quote", func() {
		updateQuote()
	})

	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	buttonContainer := container.NewHBox(newQuoteButton, quitButton)

	layout := container.NewBorder(
		nil,
		buttonContainer,
		nil,
		nil,
		container.NewVBox(
			quoteTextLabel,
			quoteAuthorLabel,
		),
	)

	myWindow.SetContent(layout)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()

}

func getQuote() Quote {
	var quotes []Quote
	if err := json.Unmarshal(quotesData, &quotes); err != nil {
		return Quote{
			Text:   "Error: unable to decode quotes file",
			Author: "",
			Tag:    "",
		}
	}

	if len(quotes) == 0 {
		return Quote{
			Text:   "Error: quotes file is empty",
			Author: "",
			Tag:    "",
		}
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomIndex := random.Intn(len(quotes))

	return quotes[randomIndex]
}
