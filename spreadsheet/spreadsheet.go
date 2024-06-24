package spreadsheet

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	Columns [][]string
}

func New(columnCount int) Spreadsheet {
	return Spreadsheet{
		Columns: make([][]string, columnCount),
	}
}

func (s *Spreadsheet) Insert(x, y int, data string) {
	extraX := x - (len(s.Columns) - 1)
	if extraX > 0 {
		s.Columns = append(s.Columns, make([][]string, extraX)...)
	}
	extraY := y - (len(s.Columns[x]) - 1)
	if extraY > 0 {
		s.Columns[x] = append(s.Columns[x], make([]string, extraY)...)
	}
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
	if x >= len(s.Columns) {
		return ""
	}
	if y >= len(s.Columns[x]) {
		return ""
	}
	return s.Columns[x][y]
}

func (s Spreadsheet) String() string {
	sb := strings.Builder{}
	cols := len(s.Columns)
	colKeys := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
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
