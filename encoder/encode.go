package encoder

import (
	"errors"
	. "json/strings"
	"reflect"
	"strconv"
	"strings"
)

var ErrUnspoortedType = errors.New("unsupported object type")

func encode(v any, sb *strings.Builder) error {
	if v == nil {
		sb.WriteString("null")
		return nil
	}

	tv := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	switch tv.Kind() {
	case reflect.Bool:
		sb.WriteString(strconv.FormatBool(rv.Bool()))

	case reflect.String:
		sb.WriteByte('"')
		sb.Write(EscapeString([]byte(rv.String())))
		sb.WriteByte('"')

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sb.WriteString(strconv.FormatInt(rv.Int(), 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		sb.WriteString(strconv.FormatUint(rv.Uint(), 10))

	case reflect.Float32, reflect.Float64:
		sb.WriteString(strconv.FormatFloat(v.(float64), 'f', -1, 64))

	case reflect.Array, reflect.Slice:
		sb.WriteByte('[')
		for i := 0; i < rv.Len(); i++ {
			if i > 0 {
				sb.WriteByte(',')
			}
			if err := encode(rv.Index(i).Interface(), sb); err != nil {
				return err
			}
		}
		sb.WriteByte(']')

	case reflect.Struct:
		numFields := rv.NumField()
		sb.WriteByte('{')
		for i := 0; i < numFields; i++ {
			if i > 0 {
				sb.WriteByte(',')
			}
			field := rv.Field(i)
			fieldName := rv.Type().Field(i).Tag.Get("json")
			if fieldName == "" {
				fieldName = CamelCaseToSnakeCase(rv.Type().Field(i).Name)
			}

			sb.WriteByte('"')
			sb.Write(EscapeString([]byte(fieldName)))
			sb.WriteByte('"')
			sb.WriteByte(':')

			if err := encode(field.Interface(), sb); err != nil {
				return err
			}
		}
		sb.WriteByte('}')
	}

	return nil
}

func Encode(v any) (string, error) {
	sb := strings.Builder{}
	err := encode(v, &sb)
	return sb.String(), err
}
