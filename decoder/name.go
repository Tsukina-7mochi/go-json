package decoder

import "strings"

func snakeCaseToUpperCamelCase(name string) string {
	toUpperCase := true
	sb := strings.Builder{}

	for _, c := range name {
		if c == '_' {
			toUpperCase = true
			continue
		}

		if toUpperCase {
			sb.WriteString(strings.ToUpper(string(c)))
		} else {
			sb.WriteRune(c)
		}
		toUpperCase = false
	}

	return sb.String()
}
