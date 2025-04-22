package main

import (
	_ "embed"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/simpleittools/fyne-quote/pkg/assets"
	"math/rand"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
}

type CustomTheme struct {
	fyne.Theme
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{Theme: theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Daily Quote")

	quoteTextLabel := widget.NewLabel("Quote:")
	quoteText := widget.NewMultiLineEntry()
	quoteText.Wrapping = fyne.TextWrapWord
	quoteText.Disable()

	quoteAuthorLabel := widget.NewLabel("Author:")
	quoteAuthor := widget.NewMultiLineEntry()
	quoteAuthor.Wrapping = fyne.TextWrapWord
	quoteAuthor.Disable()

	quoteCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		myWindow.Clipboard().SetContent(quoteText.Text)
	})

	authorCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		myWindow.Clipboard().SetContent(quoteAuthor.Text)
	})

	updateQuote := func() {
		quote := getQuote()
		quoteText.SetText(quote.Text)
		quoteAuthor.SetText(quote.Author)
	}

	updateQuote()

	newQuoteButton := widget.NewButton("New Quote", func() {
		updateQuote()
	})

	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	//contentGrid := container.New(layout.NewFormLayout(),
	//	quoteTextLabel, quoteText, quoteCopyButton,
	//	quoteAuthorLabel, quoteAuthor, authorCopyButton,
	//)
	quoteRow := container.New(layout.NewBorderLayout(nil, nil, quoteTextLabel, quoteCopyButton),
		quoteTextLabel, quoteCopyButton, quoteText)

	authorRow := container.New(layout.NewBorderLayout(nil, nil, quoteAuthorLabel, authorCopyButton),
		quoteAuthorLabel, authorCopyButton, quoteAuthor)

	contentArea := container.NewVBox(quoteRow, authorRow)

	buttonContainer := container.NewHBox(newQuoteButton, quitButton)

	pageLayout := container.NewBorder(
		nil,
		buttonContainer,
		nil,
		nil,
		contentArea,
	)

	myWindow.SetContent(pageLayout)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()

}

func getQuote() Quote {
	var quotes []Quote
	if err := json.Unmarshal(assets.GetQuotesData(), &quotes); err != nil {
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
