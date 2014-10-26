package logger

type Attrs map[string]interface{}

func SplitAttrs(v ...interface{}) ([]interface{}, *Attrs) {
	if len(v) == 0 {
		return v, nil
	}

	attrs, ok := v[len(v) -1].(Attrs)

	if !ok {
		return v, nil
	}

	v = v[:len(v) - 1]
	return v, &attrs
}
