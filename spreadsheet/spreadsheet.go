package spreadsheet

import (
	"strconv"
	"strings"
)

var (
	colKeys = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func extendSlice[T any](s []T, count int) []T {
	extend := count - (len(s) - 1)
	if extend > 0 {
		return append(s, make([]T, extend)...)
	} else {
		return s
	}
}

type Spreadsheet struct {
	Columns [][]string
}

func New(columnCount int) Spreadsheet {
	return Spreadsheet{
		Columns: make([][]string, columnCount),
	}
}

func (s *Spreadsheet) Insert(x, y int, data string) {
	s.Columns = extendSlice(s.Columns, x)
	s.Columns[x] = extendSlice(s.Columns[x], y)
	s.Columns[x][y] = data
}

func (s Spreadsheet) RowCount() int {
	rows := 0
	for x := range s.Columns {
		rows = max(rows, len(s.Columns[x]))
	}
	return rows
}

func (s Spreadsheet) Cell(x, y int) string {
	if x >= len(s.Columns) || y >= len(s.Columns[x]) {
		return ""
	}
	return s.Columns[x][y]
}

func (s Spreadsheet) String() string {
	sb := strings.Builder{}
	cols := len(s.Columns)
	sb.WriteRune('\t')
	for x := range cols {
		sb.WriteRune(colKeys[x])
		if x < cols-1 {
			sb.WriteRune('\t')
		}
	}
	sb.WriteRune('\n')
	rows := s.RowCount()
	for y := range rows {
		sb.WriteString(strconv.Itoa(y + 1))
		sb.WriteRune('\t')
		for x := range cols {
			sb.WriteString(s.Cell(x, y))
			if x < cols-1 {
				sb.WriteRune('\t')
			}
		}
		if y < rows-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}
