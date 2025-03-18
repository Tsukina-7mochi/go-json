package strings

import "errors"

func EscapeString(str []byte) []byte {
	result := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '"' || c == '\\' || c == '/' || c == '\b' || c == '\f' || c == '\n' || c == '\r' || c == '\t' {
			result = append(result, '\\')
		}

		result = append(result, c)
	}

	return result
}

func UnescapeString(str []byte) ([]byte, error) {
	result := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if i+1 >= len(str) {
				return nil, errors.New("Unexpected end of string")
			}
			i += 1

			switch str[i] {
			case '"':
				result = append(result, '"')
			case '\\':
				result = append(result, '\\')
			case '/':
				result = append(result, '/')
			case 'b':
				result = append(result, '\b')
			case 'f':
				result = append(result, '\f')
			case 'n':
				result = append(result, '\n')
			case 'r':
				result = append(result, '\r')
			case 't':
				result = append(result, '\t')
			case 'u':
				return nil, errors.New("Unicode literal is not supported")
			}
		} else {
			result = append(result, str[i])
		}
	}

	return result, nil
}
