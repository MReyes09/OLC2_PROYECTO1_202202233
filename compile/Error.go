package compile

import (
	"html/template"
	"os"
)

type ErrorReport struct {
	No      int
	Mensaje string
	Linea   int
	Columna int
	Tipo    string
}

func GenerarReporteHTML(errores []ErrorReport, nombreArchivo string) error {
	const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<title>Reporte de Errores</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/5/w3.css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Montserrat">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body,h1,h2,h3,h4,h5,h6 {font-family: "Lato", sans-serif}
.w3-bar,h1,button {font-family: "Montserrat", sans-serif}
.fa-anchor,.fa-coffee {font-size:200px}
</style>
</head>
<body>

<!-- Navbar -->
<div class="w3-top">
  <div class="w3-bar w3-red w3-card w3-left-align w3-large">
    <a class="w3-bar-item w3-button w3-hide-medium w3-hide-large w3-right w3-padding-large w3-hover-white w3-large w3-red" href="javascript:void(0);" onclick="myFunction()" title="Toggle Navigation Menu"><i class="fa fa-bars"></i></a>
    <a href="#" class="w3-bar-item w3-button w3-padding-large w3-white">Home</a>
  </div>

  <!-- Navbar on small screens -->
  <div id="navDemo" class="w3-bar-block w3-white w3-hide w3-hide-large w3-hide-medium w3-large">
  </div>
</div>

<!-- Header -->
<header class="w3-container w3-red w3-center" style="padding:128px 16px">
  <h1 class="w3-margin w3-jumbo">REPORTE DE ERRORES</h1>
  <p class="w3-xlarge">Compilador OLC2</p>
</header>

<!-- Tabla de Errores -->
<div class="w3-row-padding w3-padding-64 w3-container">
  <div class="w3-content">
    <div class="w3-twothird">
      <h2>Listado de Errores</h2>
      <div class="w3-responsive">
        <table class="w3-table-all w3-hoverable">
          <thead>
            <tr class="w3-red">
              <th>No.</th>
              <th>Mensaje</th>
              <th>Línea</th>
              <th>Columna</th>
              <th>Tipo</th>
            </tr>
          </thead>
          <tbody>
            {{range .}}
              <tr>
                <td>{{.No}}</td>
                <td>{{.Mensaje}}</td>
                <td>{{.Linea}}</td>
                <td>{{.Columna}}</td>
                <td>{{.Tipo}}</td>
              </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>
    <div class="w3-third w3-center">
      <i class="fa fa-bug w3-padding-64 w3-text-red"></i>
    </div>
  </div>
</div>

<!-- Footer -->
<footer class="w3-container w3-black w3-center w3-opacity w3-padding-64">
    <h1 class="w3-margin w3-xlarge">Compilador OLC2</h1>
</footer>

<script>
// Used to toggle the menu on small screens when clicking on the menu button
function myFunction() {
  var x = document.getElementById("navDemo");
  if (x.className.indexOf("w3-show") == -1) {
    x.className += " w3-show";
  } else { 
    x.className = x.className.replace(" w3-show", "");
  }
}
</script>

</body>
</html>
`

	tmpl, err := template.New("reporte").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, errores)
}
