package snakecase

import (
	"unicode"
)

// ToSnake convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func SnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		// Check for out[len(out)-1] is to prevent a space followed by upper-case from becoming two underscores in a row
		// e.g. "Hello World" would become hello__world
		// However this will prevent two underscores in a row, which is a weird edge-case anyway...?
		if i > 0 && unicode.IsUpper(runes[i]) && out[len(out)-1] != '_' && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		// Convert spaces to underscores
		if runes[i] == ' ' {
			out = append(out, '_')
		} else {
			out = append(out, unicode.ToLower(runes[i]))
		}
	}

	return string(out)
}
