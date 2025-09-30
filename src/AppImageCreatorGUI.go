package main

import (
	"os"
	"fmt"
	"time"
	"os/exec"
	"strings"
	"image/color"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()
	window := app.NewWindow("CreateAppImage")
	var (
		root                   *fyne.Container
		titleLabel             = canvas.NewText("CreateAppImage", nil)
		nameEntry              = widget.NewEntry()
		chooseIconButton       *widget.Button
		chooseMainBinaryButton *widget.Button
		guiButton              *widget.Button
		terminalButton         *widget.Button
		compileButton          *widget.Button

		icon            string
		mainBinary      string
		filled          int
		terminalApp     bool
	)
	window.SetFixedSize(true)
	titleLabel.TextSize = 30
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	titleLabel.Alignment = fyne.TextAlignCenter
	nameEntry.SetPlaceHolder("Enter Name for AppImage")
	chooseIconButton = widget.NewButton("Select Icon", func() {
		chooseIconWindow := app.NewWindow("Choose Icon")
		chooseIconWindow.SetFixedSize(true)
		chooseIconDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, _ error) {
			if reader == nil {return}
			filled++
			icon = reader.URI().Path()
			chooseIconButton.SetText(fmt.Sprintf("Change Icon (Current: %v)", filepath.Base(icon)))
		}, chooseIconWindow)
		chooseIconDialog.Resize(fyne.NewSize(600,400))
		chooseIconDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
		chooseIconDialog.SetOnClosed(func() {chooseIconWindow.Close()})
		chooseIconDialog.Show()
		chooseIconWindow.Resize(fyne.NewSize(600,400))
		chooseIconWindow.Show()
		fmt.Println(icon)
	})
	chooseMainBinaryButton = widget.NewButton("Select Main Binary", func() {
		chooseMainBinaryWindow := app.NewWindow("Choose Main Binary")
		chooseMainBinaryWindow.SetFixedSize(true)
		chooseMainBinaryDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, _ error) {
			if reader == nil {return}
			filled++
			mainBinary = reader.URI().Path()
			chooseMainBinaryButton.SetText(fmt.Sprintf("Change Binary (Current: %v)", filepath.Base(mainBinary)))
		}, chooseMainBinaryWindow)
		chooseMainBinaryDialog.Resize(fyne.NewSize(600,400))
		chooseMainBinaryDialog.SetOnClosed(func() {chooseMainBinaryWindow.Close()})
		chooseMainBinaryDialog.Show()
		chooseMainBinaryWindow.Resize(fyne.NewSize(600,400))
		chooseMainBinaryWindow.Show()
	})
	guiButton = widget.NewButton("GUI App", func() {
		if !(terminalButton.Importance == widget.HighImportance) {filled++}
		terminalApp = false
		terminalButton.Importance = widget.MediumImportance
		terminalButton.Refresh()
		guiButton.Importance = widget.HighImportance
		guiButton.Refresh()
	})
	terminalButton = widget.NewButton("Terminal App", func() {
		if !(guiButton.Importance == widget.HighImportance) {filled++}
		terminalApp = true
		terminalButton.Importance = widget.HighImportance
		terminalButton.Refresh()
		guiButton.Importance = widget.MediumImportance
		guiButton.Refresh()
	})
	compileButton = widget.NewButton("Compile", func() {
		var(
			statusText = canvas.NewText("", nil)
			errorText = canvas.NewText("Please fill out all fields!", color.RGBA{255, 0, 0, 255})
			compiled bool
			lastChildIsText bool
			counter int
			runCommand *exec.Cmd
		)
		statusText.Alignment = fyne.TextAlignCenter
		errorText.Alignment = fyne.TextAlignCenter
		switch root.Objects[len(root.Objects)-1].(type) {
		case *canvas.Text:
			lastChildIsText = true
		}
		if strings.TrimSpace(nameEntry.Text) == "" || filled<3 { 
			if !lastChildIsText {
				window.Resize(fyne.NewSize(450, 295))
				root.Add(errorText)
			}
			return
		}
		if lastChildIsText {root.Remove(root.Objects[len(root.Objects)-1])}
		window.Resize(fyne.NewSize(450, 295))
		root.Add(statusText)
		command := []string{os.Getenv("APPDIR") + "/AppImageCreatorCLI", "-n:"+nameEntry.Text, "-i:"+icon, "-m:"+mainBinary}
		if terminalApp {command = append(command, "-t")}
		go func() {
			for !compiled {
				switch counter % 3 {
				case 0:
					statusText.Text = "Compiling."
				case 1:
					statusText.Text = "Compiling.."
				case 2:
					statusText.Text = "Compiling..."
				}
				counter++
				fyne.Do(func() {statusText.Refresh()})
				time.Sleep(300 * time.Millisecond)
			}
		}()
		go func() {
			runCommand = exec.Command(command[0], command[1:]...)
			err := runCommand.Run()
			compiled = true
			if err == nil {
				statusText.Text = "Compiled Successfully!"
			} else {
				fmt.Println(err)
				statusText.Text = "Compile Error!"
				statusText.Color = color.RGBA{255, 0, 0, 255}
			}
			fyne.Do(func() {statusText.Refresh()})
		}()
		filled = 0
		nameEntry.SetText("")
		chooseIconButton.SetText("Select Icon")
		chooseMainBinaryButton.SetText("Select Main Binary")
		guiButton.Importance = widget.MediumImportance
		terminalButton.Importance = widget.MediumImportance
		guiButton.Refresh()
		terminalButton.Refresh()
	})
	root = container.NewVBox(
		titleLabel,
		layout.NewSpacer(),
		nameEntry, chooseIconButton,
		chooseMainBinaryButton,
		container.NewGridWithColumns(2,
			guiButton,
			terminalButton,
		),
		layout.NewSpacer(),
		compileButton,
	)
	window.SetContent(root)
	window.Resize(fyne.NewSize(450, 275))
	window.CenterOnScreen()
	window.ShowAndRun()
}