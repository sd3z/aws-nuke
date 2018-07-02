package types

import (
	"fmt"
	"strings"
)

type Properties map[string]string

func NewProperties() Properties {
	return make(Properties)
}

func (p Properties) String() string {
	parts := []string{}
	for k, v := range p {
		parts = append(parts, fmt.Sprintf(`%s: "%v"`, k, v))
	}

	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}

func (p Properties) Set(key string, value interface{}) Properties {
	if value == nil {
		return p
	}

	switch v := value.(type) {
	case *string:
		p[key] = *v
	case []byte:
		p[key] = string(v)
	case *bool:
		p[key] = fmt.Sprint(*v)
	case *int64:
		p[key] = fmt.Sprint(*v)
	case *int:
		p[key] = fmt.Sprint(*v)
	default:
		// Fallback to Stringer interface. This produces gibberish on pointers,
		// but is the only way to avoid reflection.
		p[key] = fmt.Sprint(value)
	}

	return p
}

func (p Properties) Get(key string) string {
	value, ok := p[key]
	if !ok {
		return ""
	}

	return value
}

func (p Properties) Equals(o Properties) bool {
	if p == nil && o == nil {
		return true
	}

	if p == nil || o == nil {
		return false
	}

	if len(p) != len(o) {
		return false
	}

	for k, pv := range p {
		ov, ok := o[k]
		if !ok {
			return false
		}

		if pv != ov {
			return false
		}
	}

	return true
}
