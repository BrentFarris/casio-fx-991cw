package decoder

import (
	"casiofx991cw/parser"
	"casiofx991cw/query"
	"casiofx991cw/spreadsheet"
	"errors"
	"math/bits"
	"slices"
	"strconv"
	"strings"
)

type decodedTableColumn struct {
	BitField [12]uint8
}

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
	cols := make([]decodedTableColumn, 5)
	dataCount := 0
	for i := range 5 {
		// This is a bit-field (4-bits), ever 4 rows (little-endian)
		for j, r := range t[0:12] {
			val, _ := strconv.ParseInt(string(r), 16, 8)
			cols[i].BitField[j] = bits.Reverse8(uint8(val)) >> 4
			dataCount += bits.OnesCount(uint(val))
		}
		t = t[12:]
	}
	if len(t) < dataCount*9 {
		return "", errors.New("invalid payload data for the table")
	}
	sheet := spreadsheet.New(len(cols))
	for x := range cols {
		y := 0
		for i := range cols[x].BitField {
			for j := range 4 {
				if (cols[x].BitField[i] & uint8(1<<j)) != 0 {
					cell, _ := parser.ReadNumber(t[:6], t[7:9])
					sheet.Insert(x, y, cell)
					t = t[9:]
				}
				y++
			}
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
