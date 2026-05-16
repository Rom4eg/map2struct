package main

import "reflect"

func Map2Struct(m map[string]interface{}, s interface{}) error {
	if len(m) == 0 {
		return ErrEmptyMap
	}

	sv := reflect.ValueOf(s)
	if sv.Kind() != reflect.Pointer {
		return ErrMustBePointerStruct
	}

	tags := parseTags(s)
	if len(tags) == 0 {
		return nil
	}

	for _, t := range tags {
		v, ok := m[t.Name]
		if !ok {
			continue
		}

		el := sv.Elem()
		fname := el.FieldByName(t.FName)
		if !fname.CanSet() {
			continue
		}

		switch t.Type {
		case M2SString:
			if _vt, _ok := v.(string); _ok {
				fname.SetString(_vt)
			}
		case M2SSInt:
			if _vt, _ok := v.(int); _ok {
				fname.SetInt(int64(_vt))
			}
		case M2SSFloat:
			if _vt, _ok := v.(float64); _ok {
				fname.SetFloat(_vt)
			}
		case M2SList:
			fname.Set(reflect.ValueOf(v))
		case M2SBool:
			if _vt, _ok := v.(bool); _ok {
				fname.SetBool(_vt)
			}
		}
	}
	return nil
}
