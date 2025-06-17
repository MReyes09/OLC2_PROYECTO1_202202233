package compile

import (
	"html/template"
	"os"
)

type Simbolo struct {
	Nombre string
	Tipo   string
	Valor  interface{}
	Ambito string
	Linea  int
	// Agrega más campos si lo necesitas
}

type TablaSimbolos struct {
	Simbolos map[string]Simbolo // Clave: nombre del símbolo
	// Si quieres manejar múltiples ámbitos, puedes usar una lista de mapas o una estructura más compleja
}

func NewTablaSimbolos() *TablaSimbolos {
	return &TablaSimbolos{
		Simbolos: make(map[string]Simbolo),
	}
}

func (ts *TablaSimbolos) Insertar(nombre, tipo, ambito string, linea int, valor interface{}) {
	ts.Simbolos[nombre] = Simbolo{
		Nombre: nombre,
		Tipo:   tipo,
		Valor:  valor,
		Ambito: ambito,
		Linea:  linea,
	}
}

func (ts *TablaSimbolos) Buscar(nombre string) (Simbolo, bool) {
	simbolo, existe := ts.Simbolos[nombre]
	return simbolo, existe
}

func (ts *TablaSimbolos) GenerarReporteHTML(ruta string) error {
	// Define la plantilla HTML con W3.CSS
	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <title>Tabla de Símbolos</title>
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
        .symbol-table {width: 100%; margin: 20px 0;}
        .symbol-table th {background-color: #2196f3; color: white;}
        .symbol-table tr:nth-child(even) {background-color: #f2f2f2;}
    </style>
</head>
<body>

<!-- Navbar -->
<div class="w3-top">
  <div class="w3-bar w3-blue w3-card w3-left-align w3-large">
    <a class="w3-bar-item w3-button w3-hide-medium w3-hide-large w3-right w3-padding-large w3-hover-white w3-large w3-blue" href="javascript:void(0);" onclick="myFunction()" title="Toggle Navigation Menu"><i class="fa fa-bars"></i></a>
    <a href="#" class="w3-bar-item w3-button w3-padding-large w3-white">Tabla de Símbolos</a>
  </div>

  <!-- Navbar on small screens -->
  <div id="navDemo" class="w3-bar-block w3-white w3-hide w3-hide-large w3-hide-medium w3-large">
    <a href="#" class="w3-bar-item w3-button w3-padding-large">Tabla de Símbolos</a>
  </div>
</div>

<!-- Header -->
<header class="w3-container w3-blue w3-center" style="padding:128px 16px">
<img src="../img/logo.png" alt="Logo de la empresa">

  <h1 class="w3-margin w3-jumbo">Tabla de Símbolos</h1>
  <p class="w3-xlarge">Reporte de símbolos del compilador</p>
</header>

<!-- First Grid -->
<div class="w3-row-padding w3-padding-64 w3-container">
  <div class="w3-content">
    <div class="w3-twothird">
      <h1>Símbolos encontrados</h1>
      <table class="w3-table-all symbol-table">
        <tr>
          <th>Nombre</th>
          <th>Tipo</th>
          <th>Valor</th>
          <th>Ámbito</th>
          <th>Línea</th>
        </tr>
        {{range .Simbolos}}
        <tr>
          <td>{{.Nombre}}</td>
          <td>{{.Tipo}}</td>
          <td>{{.Valor}}</td>
          <td>{{.Ambito}}</td>
          <td>{{.Linea}}</td>
        </tr>
        {{end}}
      </table>
    </div>
    <div class="w3-third w3-center">
      <i class="fa fa-code w3-padding-64 w3-text-blue"></i>
    </div>
  </div>
</div>


<div class="w3-container w3-black w3-center w3-opacity w3-padding-64">
    <h1 class="w3-margin w3-xlarge">Compiladores 2</h1>
</div>



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
</html>`

	// Crea el archivo
	file, err := os.Create(ruta)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parsea y ejecuta la plantilla
	t := template.Must(template.New("tabla").Parse(tmpl))
	return t.Execute(file, ts)
}
