package controller

import (
	"strings"
)

// FormatAntlrTree takes a flat ANTLR parse tree string and formats it with indentation.
// It assumes the input string uses parentheses to denote tree structure,
// like `(parent (child1) (child2 (grandchild)))`.
func FormatAntlrTree(treeString string) string {
	var formattedTree strings.Builder
	indentationLevel := 0

	// First, normalize the input string to ensure consistent spacing around structural elements
	// and to handle string literals correctly.
	normalizedString := ""
	inQuote := false // Flag to track if we're inside a double-quoted string literal

	for _, r := range treeString {
		if r == '"' {
			inQuote = !inQuote // Toggle inQuote flag
			normalizedString += string(r)
		} else if inQuote {
			normalizedString += string(r) // Keep all characters as is within quotes
		} else {
			switch r {
			case '(', ')', '{', '}': // Ensure spaces around these structural characters
				normalizedString += " " + string(r) + " "
			default:
				normalizedString += string(r)
			}
		}
	}

	// Clean up multiple spaces that might have been introduced
	normalizedString = strings.Join(strings.Fields(normalizedString), " ")
	normalizedString = strings.TrimSpace(normalizedString)

	// Now, iterate through the normalized string, token by token (separated by space)
	// to apply indentation.
	tokens := strings.Fields(normalizedString) // Split by any whitespace

	for i, token := range tokens {
		switch token {
		case "(":
			formattedTree.WriteString(token)
			indentationLevel++
			formattedTree.WriteRune('\n')
			for j := 0; j < indentationLevel; j++ {
				formattedTree.WriteRune('\t')
			}
		case ")":
			indentationLevel--
			// Add a newline and tabs before closing parenthesis if content precedes it on the same line
			if formattedTree.Len() > 0 && formattedTree.String()[formattedTree.Len()-1] != '\n' && formattedTree.String()[formattedTree.Len()-1] != '\t' {
				formattedTree.WriteRune('\n')
			}
			for j := 0; j < indentationLevel; j++ {
				formattedTree.WriteRune('\t')
			}
			formattedTree.WriteString(token)
			// Add a space after ')' if it's not the end of the string
			if i < len(tokens)-1 {
				formattedTree.WriteRune(' ')
			}
		case "{":
			// Handle opening curly brace similar to '(' for indentation
			if formattedTree.Len() > 0 && formattedTree.String()[formattedTree.Len()-1] != '\n' && formattedTree.String()[formattedTree.Len()-1] != '\t' {
				formattedTree.WriteRune(' ')
			}
			formattedTree.WriteString(token)
			indentationLevel++
			formattedTree.WriteRune('\n')
			for j := 0; j < indentationLevel; j++ {
				formattedTree.WriteRune('\t')
			}
		case "}":
			// Handle closing curly brace similar to ')' for indentation
			indentationLevel--
			if formattedTree.Len() > 0 && formattedTree.String()[formattedTree.Len()-1] != '\n' && formattedTree.String()[formattedTree.Len()-1] != '\t' {
				formattedTree.WriteRune('\n')
			}
			for j := 0; j < indentationLevel; j++ {
				formattedTree.WriteRune('\t')
			}
			formattedTree.WriteString(token)
			// Add a space after '}' if it's not the end of the string
			if i < len(tokens)-1 {
				formattedTree.WriteRune(' ')
			}
		default:
			// For regular tokens, add a space before if it's not the first token in a line
			// and not immediately after an opening parenthesis/brace.
			if formattedTree.Len() > 0 && formattedTree.String()[formattedTree.Len()-1] != '(' && formattedTree.String()[formattedTree.Len()-1] != '{' && formattedTree.String()[formattedTree.Len()-1] != '\n' && formattedTree.String()[formattedTree.Len()-1] != '\t' {
				formattedTree.WriteRune(' ')
			}
			formattedTree.WriteString(token)
		}
	}

	// Final trim to clean up any leading/trailing whitespace
	return strings.TrimSpace(formattedTree.String())
}
