package primitvos

func StringToByte(s string) []byte {
	cadena := []byte(s)
	cadena = append(cadena, 0) // Agregar byte nulo al final (terminador)
	return cadena
}
