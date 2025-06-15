// controller/dot_generator.go
package controller

import (
	"fmt"
	"os"
	"strings"
)

type nodeInfo struct {
	id    int
	label string
	depth int
}

func GenerateDotFromFormattedTreeString(formattedTreeString, dotFilePath string) error {
	lines := strings.Split(formattedTreeString, "\n")
	stack := []nodeInfo{}
	nodeCounter := 0
	dotContent := strings.Builder{}
	dotContent.WriteString("digraph G {\n")
	dotContent.WriteString("  node [shape=box];\n")

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		depth := 0
		for _, r := range line {
			if r == '\t' {
				depth++
			} else {
				break
			}
		}

		content := strings.TrimSpace(line)
		currentNode := nodeInfo{
			id:    nodeCounter,
			label: escapeForDot(content),
			depth: depth,
		}
		nodeCounter++

		// Buscar padre en el stack (mismo nivel o menor)
		for len(stack) > 0 && stack[len(stack)-1].depth >= depth {
			stack = stack[:len(stack)-1] // Pop hasta encontrar el padre
		}

		// Conectar con padre si existe
		if len(stack) > 0 {
			parent := stack[len(stack)-1]
			dotContent.WriteString(fmt.Sprintf("  node%d -> node%d;\n", parent.id, currentNode.id))
		}

		// Agregar nodo al DOT
		dotContent.WriteString(fmt.Sprintf("  node%d [label=\"%s\"];\n", currentNode.id, currentNode.label))

		// Manejar nodos de cierre especiales (solo paréntesis/llaves)
		if content == ")" || content == "}" {
			// Nodos de cierre no se apilan
			if i < len(lines)-1 {
				stack = stack[:len(stack)-1] // Cierra el nodo padre
			}
		} else {
			stack = append(stack, currentNode) // Apilar nodos normales
		}
	}

	dotContent.WriteString("}\n")
	return os.WriteFile(dotFilePath, []byte(dotContent.String()), 0644)
}

func escapeForDot1(s string) string {
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\\", "\\\\")
	return s
}

func escapeForDot(s string) string {
	// Manejar secuencias de escape comunes
	s = strings.ReplaceAll(s, "\\", "\\\\") // Escapar barras invertidas primero
	s = strings.ReplaceAll(s, "\"", "\\\"") // Luego escapar comillas
	s = strings.ReplaceAll(s, "\n", "\\n")  // Escapar saltos de línea
	s = strings.ReplaceAll(s, "\t", "\\t")  // Escapar tabs
	return s
}
