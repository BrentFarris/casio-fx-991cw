package decoder

import (
	"casiofx991cw/query"
	"errors"
	"slices"
	"strings"
)

type Encoding struct {
	Key           string
	Transform     string
	TransformFunc func(in string, out *strings.Builder) (int, error)
}

type Decoder struct {
	encodings []Encoding
}

func New() Decoder {
	return Decoder{
		encodings: make([]Encoding, 0),
	}
}

func (d *Decoder) AddEncoding(encoding Encoding) {
	d.encodings = append(d.encodings, encoding)
	slices.SortFunc(d.encodings, func(a, b Encoding) int {
		return len(b.Key) - len(a.Key)
	})
}

func (d *Decoder) Decode(query query.Query) (string, error) {
	sb := strings.Builder{}
	expr := query.E
	for expr != "" {
		read := false
		for i := range d.encodings {
			if strings.HasPrefix(expr, d.encodings[i].Key) {
				expr = expr[len(d.encodings[i].Key):]
				if d.encodings[i].Transform != "" {
					sb.WriteString(d.encodings[i].Transform)
				} else if d.encodings[i].TransformFunc != nil {
					read, err := d.encodings[i].TransformFunc(expr, &sb)
					if err != nil {
						return sb.String(), err
					}
					expr = expr[read:]
				} else {
					return sb.String(), errors.New("malformed encoding " + d.encodings[i].Key)
				}
				read = true
				break
			}
		}
		if !read {
			return sb.String(), errors.New("Failed to find encoding for sequence " + expr)
		}
	}
	sb.WriteRune('=')
	sb.WriteString(query.Answer)
	return sb.String(), nil
}
