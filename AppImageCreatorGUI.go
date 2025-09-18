package main

import (
	"fmt"
	"image/color"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	notFilled := false
	icon := ""
	mainBinary := ""
	app := app.New()
	window := app.NewWindow("CreateAppImage")
	var root *fyne.Container
	var chooseIconButton *widget.Button
	var chooseMainBinaryButton *widget.Button
	titleLabel := canvas.NewText("CreateAppImage", nil)
	nameEntry := widget.NewEntry()
	window.SetFixedSize(true)
	chooseIconButton = widget.NewButton("Select Icon", func() {
		chooseIconWindow := app.NewWindow("Choose Icon")
		chooseIconWindow.SetFixedSize(true)

		chooseIconDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, _ error) {
			if reader == nil {
				notFilled = true
				return
			}
			icon = reader.URI().Path()
			chooseIconButton.SetText(fmt.Sprintf("Change Icon (Current: %v)", filepath.Base(icon)))
		}, chooseIconWindow)
		chooseIconDialog.Resize(fyne.NewSize(600,400))
		chooseIconDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
		chooseIconDialog.SetOnClosed(func(){chooseIconWindow.Close()})
		chooseIconDialog.Show()

		chooseIconWindow.Resize(fyne.NewSize(600,400))
		chooseIconWindow.Show()
		fmt.Println(icon)
	})
	chooseMainBinaryButton = widget.NewButton("Select Main Binary", func() {
		chooseMainBinaryWindow := app.NewWindow("Choose Main Binary")
		chooseMainBinaryWindow.SetFixedSize(true)

		chooseMainBinaryDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, _ error) {
			if reader == nil {
				notFilled = true
				return
			}
			mainBinary = reader.URI().Path()
			chooseMainBinaryButton.SetText(fmt.Sprintf("Change Binary (Current: %v)", filepath.Base(mainBinary)))
		}, chooseMainBinaryWindow)
		chooseMainBinaryDialog.Resize(fyne.NewSize(600,400))
		chooseMainBinaryDialog.SetFilter(storage.NewExtensionFileFilter([]string{""}))
		chooseMainBinaryDialog.SetOnClosed(func(){chooseMainBinaryWindow.Close()})
		chooseMainBinaryDialog.Show()

		chooseMainBinaryWindow.Resize(fyne.NewSize(600,400))
		chooseMainBinaryWindow.Show()
	})
		submitButton := widget.NewButton("Submit", func() {
			if strings.TrimSpace(nameEntry.Text) == "" || notFilled {
				root.Add(canvas.NewText("Please fill out all fields!", color.RGBA{255, 0, 0, 255}))
				return
			}
			_, lastVarIsText := root.Objects[len(root.Objects)-1].(*canvas.Text)
			if lastVarIsText {
				root.Remove(root.Objects[len(root.Objects)-1])
			}
			exec.Command("./AppImageCreatorCLI", "-n:"+nameEntry.Text, "-i:"+icon, "-m:"+mainBinary)
		})

	titleLabel.TextSize = 30
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	titleLabel.Alignment = fyne.TextAlignCenter
	nameEntry.SetPlaceHolder("Enter Name for AppImage")

	root = container.NewVBox(titleLabel, layout.NewSpacer(), nameEntry, chooseIconButton, chooseMainBinaryButton, layout.NewSpacer(), submitButton)
	window.SetContent(root)
	window.Resize(fyne.NewSize(400, 250))
	window.CenterOnScreen()
	window.ShowAndRun()
}