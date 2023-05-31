package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var data = []string{"a", "string", "list"}

func main() {
	ecouteApp := app.New()
	appWindow := ecouteApp.NewWindow("Ecoute")

	appWindow.SetContent(createContent())
	appWindow.Resize(fyne.NewSize(1000, 600))
	appWindow.SetMaster()
	appWindow.CenterOnScreen()
	appWindow.ShowAndRun()
}

func audioDeviceChange(device string) {
	log.Println("Audio Device Changed!")
}

func createContent() fyne.CanvasObject {

	transcriptText := "This is a long transcript text that needs to be wrapped and centered."
	transcriptTextbox := widget.NewMultiLineEntry()
	transcriptTextbox.Resize(fyne.NewSize(300, 400))
	transcriptTextbox.TextStyle.Bold = true
	transcriptTextbox.Wrapping = fyne.TextWrapBreak
	transcriptTextbox.SetText(transcriptText)
	transcriptTextbox.Disable()

	responseText := "This is a long response text that needs to be wrapped and centered."
	responseTextbox := widget.NewMultiLineEntry()
	responseTextbox.Resize(fyne.NewSize(300, 400))
	responseTextbox.TextStyle.Bold = true
	responseTextbox.Wrapping = fyne.TextWrapBreak
	responseTextbox.SetText(responseText)
	responseTextbox.Disable()

	transcriptContainer := container.New(layout.NewGridLayout(2))
	transcriptContainer.Resize(fyne.NewSize(1000, 400))
	transcriptContainer.Add(transcriptTextbox)
	transcriptContainer.Add(responseTextbox)

	clearButton := widget.NewButton("Clear Transcript", func() {})

	freezeButton := widget.NewButton("Freeze", func() {})

	ButtonContainer := container.New(layout.NewGridLayout(2), clearButton, freezeButton)

	// Input Devices
	inputDevicesListLabel := widget.NewLabel("Select Speaker")
	inputDevicesListLabel.TextStyle.Monospace = true
	inputDevicesListLabel.Resize(fyne.NewSize(0, 30))

	inputDevicesList := widget.NewSelect(data, audioDeviceChange)

	inputDevicesContainer := container.New(layout.NewVBoxLayout(), inputDevicesListLabel, inputDevicesList)

	// Options
	updateIntervalSliderLabel := widget.NewLabel("Interval")
	updateIntervalSliderLabel.TextStyle.Monospace = true
	updateIntervalSliderLabel.Resize(fyne.NewSize(0, 30))

	updateIntervalSlider := widget.NewSlider(1, 10)
	updateIntervalSlider.Resize(fyne.NewSize(300, 20))
	updateIntervalSlider.Step = 2

	sliderContainer := container.New(layout.NewVBoxLayout(), updateIntervalSliderLabel, updateIntervalSlider)

	OptionsContainer := container.New(layout.NewGridLayout(2), inputDevicesContainer, sliderContainer)
	appContainer := container.NewVBox(transcriptContainer, layout.NewSpacer(), ButtonContainer, OptionsContainer)

	return appContainer
}
