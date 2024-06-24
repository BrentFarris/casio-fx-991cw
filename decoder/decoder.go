package decoder

import (
	"casiofx991cw/parser"
	"casiofx991cw/query"
	"casiofx991cw/spreadsheet"
	"errors"
	"fmt"
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

func (d *Decoder) decodeE(query query.Query) (string, error) {
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

func (d *Decoder) decodeTable(q query.Query) (string, error) {
	if !strings.HasPrefix(q.T, "SP") {
		return "", errors.New("invalid prefix to T, expected 'SP'")
	}
	t := q.T[2:]
	// 12 character sequences for column lengths (5 entries)
	if len(t) < 12*5 {
		return "", errors.New("table prefix data is invalid")
	}
	cols := make([]int, 5)
	for i := range 5 {
		for _, r := range t[0:12] {
			switch r {
			case '8':
				cols[i] += 1
			case 'C':
				cols[i] += 2
			case 'E':
				cols[i] += 3
			case 'F':
				cols[i] += 4
			}
		}
		t = t[12:]
	}
	sum := 0
	for i := range cols {
		sum += cols[i]
	}
	if len(t) < sum*9 {
		return "", errors.New("invalid payload data for the table")
	}
	sheet := spreadsheet.New(len(cols))
	for i := range cols {
		for j := range cols[i] {
			// The last number in the entry is the number of digits in sequence
			// TODO:  Support fractions
			cell, _ := parser.ReadNumber(t[:6], t[7:9])
			if cell == "" {
				return "", fmt.Errorf("failed to read the cell %d:%d data length", i, j)
			}
			sheet.Columns[i] = append(sheet.Columns[i], cell)
			t = t[9:]
		}
	}
	return sheet.String(), nil
}

func (d *Decoder) Decode(query query.Query) (string, error) {
	if query.IsSpreadsheet() {
		return d.decodeTable(query)
	} else {
		return d.decodeE(query)
	}
}
