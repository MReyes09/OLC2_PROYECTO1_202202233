package main

import (
	"io"
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
	myWindow.Resize(fyne.NewSize(800, 600))

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
	btnAbrirArchivo := widget.NewButton("📁 Abrir Archivo", abrirArchivo)
	btnBorrarDatos := widget.NewButton("🗑️ Borrar datos", borrarDatos)
	btnEjecutar := widget.NewButton("▶️ Ejecutar", ejecutar)

	// Crear container para los botones
	botonesContainer := container.NewHBox(
		btnAbrirArchivo,
		btnBorrarDatos,
		btnEjecutar,
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
	myWindow.ShowAndRun()
}
