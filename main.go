package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	re  = regexp.MustCompile(`(\?|&)qr{0,1}=(.*?)(?:&|$)`)
	ire = regexp.MustCompile(`I-.*?(?:(\w-|$))`)
	ure = regexp.MustCompile(`U-.*?(?:(\w-|$))`)
	mre = regexp.MustCompile(`M-.*?(?:(\w-|$))`)
	sre = regexp.MustCompile(`S-.*?(?:(\w-|$))`)
	qre = regexp.MustCompile(`Q-.*?(?:(\w-|$))`)
	ere = regexp.MustCompile(`E-.*?(?:(\w-|$))`)
)

type EquationCode = int

const (
	qNumberRange = 24
)

const (
	EquationCodeUnknown EquationCode = iota
	EquationCodeSimple
	EquationCodeFraction
)

var equationCodes = map[string]EquationCode{
	"3":       EquationCodeSimple,
	"C81D1A3": EquationCodeFraction,
}

var operations = map[string]string{
	"A63": "+",
	"A73": "-",
	"A93": "÷",
	"A83": "×",
	"2E":  ".",
}

type Query struct {
	Raw          string
	I            string
	U            string
	M            string
	S            string
	Q            string
	E            string
	Answer       string
	AnswerLength int
}

func (q *Query) clean() {
	if q.IsValid() {
		q.I = strings.TrimPrefix(q.I, "I-")
		q.U = strings.TrimPrefix(q.U, "U-")
		q.M = strings.TrimPrefix(q.M, "M-")
		q.S = strings.TrimPrefix(q.S, "S-")
		q.Q = strings.TrimPrefix(q.Q, "Q-")
		q.E = strings.TrimPrefix(q.E, "E-")
		if length, err := strconv.Atoi(strings.TrimPrefix(q.Q[qNumberRange+2:28], "0")); err == nil {
			q.AnswerLength = length + 1
		}
		num := q.Q[1 : qNumberRange+1]
		if q.AnswerLength > qNumberRange {
			// This is a decimal starting from 99 being 0.x and 98 being 0.0x
			sb := strings.Builder{}
			sb.WriteString("0.")
			for i := range 100 - qNumberRange {
				if 100-i == q.AnswerLength {
					break
				}
				sb.WriteRune('0')
			}
			// Write all the numbers up to the first occurrence of 0
			end := len(num) - 1
			for range len(num) {
				if num[end] != '0' {
					break
				}
				end--
			}
			sb.WriteString(num[:end+1])
			q.Answer = sb.String()
		} else {
			q.Answer = num[:q.AnswerLength]
		}
	}
}

func pullNumber(equ string, skip rune, terminator rune) (string, int) {
	sb := strings.Builder{}
	runes := []rune(equ)
	runesRead := len(runes)
	for i := range len(runes) {
		if i%2 == 1 {
			if runes[i] == terminator {
				runesRead = i
				break
			}
			if runes[i] != skip {
				if i < len(runes)-1 && runes[i] == '2' && runes[i+1] == 'E' {
					sb.WriteRune('.')
				} else {
					panic("unexpected format")
				}
			}
			continue
		}
		sb.WriteRune(runes[i])
	}
	return sb.String(), runesRead
}

func (q Query) equation() string {
	ec := EquationCodeUnknown
	for k, v := range equationCodes {
		if strings.HasPrefix(q.E, k) {
			ec = v
		}
	}
	switch ec {
	case EquationCodeSimple:
		sb := strings.Builder{}
		e := q.E[1:]
		num1, read1 := pullNumber(e, '3', 'A')
		e = e[read1:]
		if e == "" {
			return sb.String()
		}
		sb.WriteString(num1)
		op := operations[e[0:3]]
		e = e[3:]
		sb.WriteString(op)
		num2, read2 := pullNumber(e, '3', 'A')
		e = e[read2:]
		sb.WriteString(num2)
		return sb.String()
	case EquationCodeUnknown:
		panic("unknown equation code")
	default:
		panic("equation code not yet implemented")
	}
}

func (q Query) IsValid() bool {
	return q.I != "" && q.U != "" && q.M != "" && q.S != "" && q.Q != "" && q.E != ""
}

func (q Query) Print() string {
	return q.equation() + "=" + q.Answer
}

func pullSegment(reg *regexp.Regexp, query string) string {
	matches := reg.FindAllStringSubmatch(query, -1)
	if len(matches) == 0 {
		return ""
	}
	if len(matches[0]) < 2 {
		return ""
	}
	return strings.TrimSuffix(matches[0][0], matches[0][1])
}

func Parse(url string) Query {
	matches := re.FindAllStringSubmatch(url, -1)
	if len(matches) == 0 || len(matches[0]) < 3 {
		return Query{}
	}
	query := matches[0][2]
	q := Query{
		Raw: query,
		I:   pullSegment(ire, query),
		U:   pullSegment(ure, query),
		M:   pullSegment(mre, query),
		S:   pullSegment(sre, query),
		Q:   pullSegment(qre, query),
		E:   pullSegment(ere, query),
	}
	q.clean()
	return q
}

type Test struct {
	Expected string
	URL      string
}

func main() {
	tests := []Test{
		{
			"3×9=27",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-02700000000000000000000001010000000000000000000000000000+E-33A839",
		},
		{
			"4×4=16",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01600000000000000000000001010000000000000000000000000000+E-34A834",
		},
		{
			"2×5=10",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01000000000000000000000001010000000000000000000000000000+E-32A835",
		},
		{
			"7×6=42",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-04200000000000000000000001010000000000000000000000000000+E-37A836",
		},
		{
			"8×8=64",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-06400000000000000000000001010000000000000000000000000000+E-38A838",
		},
		{
			"=9",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000001000000000000000000000000000000+E-39",
		},
		{
			"9×9=81",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-08100000000000000000000001010000000000000000000000000000+E-39A839",
		},
		{
			"3+7=10",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01000000000000000000000001010000000000000000000000000000+E-33A637",
		},
		{
			"319+1=320",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-03200000000000000000000001020000000000000000000000000000+E-333139A631",
		},
		{
			"=3200",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-03200000000000000000000001030000000000000000000000000000+E-33323030",
		},
		{
			"=0.0009", // 9×10^-4
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000000960000000000000000000000000000+E-302E30303039",
		},
		{
			"=0.09",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000000980000000000000000000000000000+E-302E3039",
		},
		{
			"3÷4=0.75",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-07500000000000000000000000990000000000000000000000000000+E-33A934",
		},
		{
			"=123456789123456789",
			//"123456789123456789=1.234567891x10^17",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01234567891234567890000001170000000000000000000000000000+E-313233343536373839313233343536373839",
		},
		{
			"35÷555=0.063063063063063063063063",
			"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-06306306306306306306306300980000000000000000000000000000+E-3335A9353535",
		},
		/*
			{
				"3+6(5+4)=57",
				"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-05700000000000000000000001010000000000000000000000000000+E-33A6366035A634D0",
			},
			{
				"3/4=0.75",
				"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-23A40000000000000000000001030000000000000000000000000000+E-C81D1A331B1A341B1E",
			},
			{
				"123/456=0.2697368421",
				"http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-241A1520000000000000000001060000000000000000000000000000+E-C81D1A3132331B1A3435361B1E",
			},
		*/
	}

	for i := range tests {
		q := Parse(tests[i].URL)
		if !q.IsValid() {
			panic("test " + strconv.Itoa(i) + " failed")
		}
		if q.Print() != tests[i].Expected {
			panic("test " + strconv.Itoa(i) + " failed; expected")
		}
		println(q.Print())
	}
}
