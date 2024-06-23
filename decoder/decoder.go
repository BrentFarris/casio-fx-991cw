package decoder

import (
	"casiofx991cw/query"
	"errors"
	"slices"
	"strings"
)

const ()

type Encoding struct {
	Key           string
	Transform     string
	TransformFunc func(in string, out *strings.Builder) int
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
		return len(a.Key) - len(b.Key)
	})
}

func (d *Decoder) Decode(query query.Query) (string, error) {
	sb := strings.Builder{}
	expr := query.E
	for expr != "" {
		for i := range d.encodings {
			if strings.HasPrefix(expr, d.encodings[i].Key) {
				expr = expr[len(d.encodings[i].Key):]
				if d.encodings[i].Transform != "" {
					sb.WriteString(d.encodings[i].Transform)
				} else if d.encodings[i].TransformFunc != nil {
					read := d.encodings[i].TransformFunc(expr, &sb)
					expr = expr[read:]
				} else {
					return "", errors.New("malformed encoding " + d.encodings[i].Key)
				}
				break
			}
		}
	}
	sb.WriteRune('=')
	sb.WriteString(query.Answer)
	return sb.String(), nil
}
