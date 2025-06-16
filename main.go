package main

import (
	"io"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	controller "OLC2CLIENTE/controller"
)

func main() {
	// Crear la aplicación
	myApp := app.New()
	myWindow := myApp.NewWindow("Compiladores 2")
	myWindow.Resize(fyne.NewSize(1400, 800))

	// Variables globales para los widgets
	var codigoEntry *widget.Entry
	var consolaEntry *widget.Entry

	// Crear las áreas de texto
	codigoEntry = widget.NewMultiLineEntry()
	codigoEntry.SetPlaceHolder("Escribe tu código esta área...")
	codigoEntry.Wrapping = fyne.TextWrapWord

	consolaEntry = widget.NewMultiLineEntry()
	consolaEntry.SetPlaceHolder("El resultado se mostrará aquí...")
	consolaEntry.Wrapping = fyne.TextWrapWord

	// Función para abrir archivo
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

	// Función para borrar datos
	borrarDatos := func() {
		codigoEntry.SetText("")
		consolaEntry.SetText("")
	}

	// Abrir Reporte
	abrirErrores := func() {
		filePath := "reports/reporte_errores.html"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default: // linux y otros
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}

	// Abrir Reporte
	abrirSimbolos := func() {
		filePath := "reports/tabla_simbolos.html"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default: // linux y otros
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}

	// Abrir Reporte
	abrirAST := func() {
		filePath := "reports/ast.pdf"
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", filePath)
		case "darwin":
			cmd = exec.Command("open", filePath)
		default: // linux y otros
			cmd = exec.Command("xdg-open", filePath)
		}
		err := cmd.Start()
		if err != nil {
			dialog.ShowError(err, myWindow)
		}
	}
	// Función para ejecutar código (simulación)
	ejecutar := func() {
		codigo := codigoEntry.Text
		if strings.TrimSpace(codigo) == "" {
			consolaEntry.SetText("Error: No hay código para ejecutar")
			return
		} else {
			output := controller.CompileCode(codigo)
			consolaEntry.SetText(output)
		}
	}

	// Crear los botones
	btnAbrirArchivo := widget.NewButton("📂 Abrir Archivo", abrirArchivo)
	btnBorrarDatos := widget.NewButton("🗑️ Borrar datos", borrarDatos)
	btnEjecutar := widget.NewButton("▶️ Ejecutar", ejecutar)
	btnReporte := widget.NewButton("⛔ Tabla de Errores", abrirErrores)
	btnSimbolos := widget.NewButton("🏷️ Tabla de Simbolos", abrirSimbolos)
	btnAST := widget.NewButton("📊 Arbol AST", abrirAST)

	// Crear container para los botones
	botonesContainer := container.NewHBox(
		btnAbrirArchivo,
		btnBorrarDatos,
		btnEjecutar,
		btnSimbolos,
		btnReporte,
		btnAST,
	)

	// Crear labels para las secciones
	labelCodigo := widget.NewLabel("📝 Código")
	labelCodigo.TextStyle.Bold = true

	labelConsola := widget.NewLabel("💻 Consola")
	labelConsola.TextStyle.Bold = true

	// Crear containers para cada sección con scroll
	codigoContainer := container.NewBorder(
		labelCodigo, nil, nil, nil,
		container.NewScroll(codigoEntry),
	)

	consolaContainer := container.NewBorder(
		labelConsola, nil, nil, nil,
		container.NewScroll(consolaEntry),
	)

	// Crear split container horizontal para código y consola
	splitContainer := container.NewHSplit(codigoContainer, consolaContainer)
	splitContainer.SetOffset(0.5) // 50% para cada lado

	// Crear el layout principal usando border layout
	content := container.NewBorder(
		botonesContainer, // top
		nil,              // bottom
		nil,              // left
		nil,              // right
		splitContainer,   // center
	)

	// Configurar y mostrar la ventana
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
