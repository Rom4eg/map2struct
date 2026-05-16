package map2struct

import (
	"reflect"
)

const TagM2S = "m2s"

func parseTags(s interface{}) []M2STag {
	actual := reflect.ValueOf(s)
	if actual.Kind() == reflect.Pointer {
		actual = actual.Elem()
	}

	if actual.Kind() != reflect.Struct {
		return nil
	}

	t := actual.Type()
	var tags []M2STag
	for i := 0; i < actual.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(TagM2S)
		if tag == "" {
			continue
		}

		switch f.Type.Kind() {
		case reflect.String:
			tags = append(tags, New(f.Name, tag, M2SString))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			tags = append(tags, New(f.Name, tag, M2SSInt))
		case reflect.Float32, reflect.Float64:
			tags = append(tags, New(f.Name, tag, M2SSFloat))
		case reflect.Slice, reflect.Array:
			tags = append(tags, New(f.Name, tag, M2SList))
		case reflect.Bool:
			tags = append(tags, New(f.Name, tag, M2SBool))
		}
	}

	return tags
}
