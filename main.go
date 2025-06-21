package main

import (
	"io"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	controller "OLC2CLIENTE/controller"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Compiladores 2")
	myWindow.Resize(fyne.NewSize(1400, 800))

	var codigoEntry *widget.Entry
	var consolaEntry *widget.Entry

	codigoEntry = widget.NewMultiLineEntry()
	codigoEntry.SetPlaceHolder("Escribe tu código esta área...")
	codigoEntry.Wrapping = fyne.TextWrapWord

	consolaEntry = widget.NewMultiLineEntry()
	consolaEntry.SetPlaceHolder("El resultado se mostrará aquí...")
	consolaEntry.Wrapping = fyne.TextWrapWord

	// Cambiar entre ver codigo interpretado y compilado
	var interpretedCode string
	var compiledCode string

	abrirArchivo := func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()
			data, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			codigoEntry.SetText(string(data))
		}, myWindow)
	}

	borrarDatos := func() {
		codigoEntry.SetText("")
		consolaEntry.SetText("")
	}

	abrirErrores := func() {
		filePath := "reports/reporte_errores.html"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default:
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}

	abrirSimbolos := func() {
		filePath := "reports/tabla_simbolos.html"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default:
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}

	abrirAST := func() {
		filePath := "reports/ast.pdf"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default:
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}

	ejecutar := func() {
		codigo := codigoEntry.Text
		if strings.TrimSpace(codigo) == "" {
			consolaEntry.SetText("Error: No hay código para ejecutar")
			return
		}
		interpretedCode, compiledCode = controller.CompileCode(codigo)
		consolaEntry.SetText(interpretedCode)
	}

	option := 0
	changeConsola := func() {
		if option == 0 {
			consolaEntry.SetText(compiledCode)
			option = 1
		} else {
			consolaEntry.SetText(interpretedCode)
			option = 0
		}

	}

	btnAbrirArchivo := widget.NewButton("📂 Abrir Archivo", abrirArchivo)
	btnBorrarDatos := widget.NewButton("🗑️ Borrar datos", borrarDatos)
	btnEjecutar := widget.NewButton("▶️ Ejecutar", ejecutar)
	btnReporte := widget.NewButton("⛔ Tabla de Errores", abrirErrores)
	btnSimbolos := widget.NewButton("🏷️ Tabla de Simbolos", abrirSimbolos)
	btnAST := widget.NewButton("📊 Arbol AST", abrirAST)
	btnCompie_Interpreted := widget.NewButton("🧾 Código Interpretado", changeConsola)

	// Imagen con tamaño fijo (200x100)
	logo := canvas.NewImageFromFile("img/logo.png")
	logo.FillMode = canvas.ImageFillStretch // Usa Stretch para forzar el tamaño exacto
	logo.SetMinSize(fyne.NewSize(200, 160)) // ancho x alto

	logoBox := container.NewCenter(logo)

	botonesContainer := container.NewHBox(
		btnAbrirArchivo,
		btnBorrarDatos,
		btnEjecutar,
		btnSimbolos,
		btnReporte,
		btnAST,
		btnCompie_Interpreted,
	)

	labelCodigo := widget.NewLabel("📝 Código")
	labelCodigo.TextStyle.Bold = true

	labelConsola := widget.NewLabel("💻 Consola")
	labelConsola.TextStyle.Bold = true

	codigoContainer := container.NewBorder(
		labelCodigo, nil, nil, nil,
		container.NewScroll(codigoEntry),
	)

	consolaContainer := container.NewBorder(
		labelConsola, nil, nil, nil,
		container.NewScroll(consolaEntry),
	)

	splitContainer := container.NewHSplit(codigoContainer, consolaContainer)
	splitContainer.SetOffset(0.5)

	logoYBotones := container.NewVBox(
		logoBox,
		botonesContainer,
	)

	content := container.NewBorder(
		logoYBotones,
		nil,
		nil,
		nil,
		splitContainer,
	)

	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
