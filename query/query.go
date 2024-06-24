package query

import (
	"casiofx991cw/parser"
	"regexp"
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
	tre = regexp.MustCompile(`T-.*?(?:(\w-|$))`)
)

const (
	qNumberRange = 24
)

type Query struct {
	Raw     string
	I       string
	U       string
	M       string
	S       string
	Q       string
	E       string
	T       string
	Answer  string
	Decimal int
}

func New(url string) Query {
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
		T:   pullSegment(tre, query),
	}
	q.clean()
	return q
}

func (q *Query) IsSpreadsheet() bool { return q.T != "" }

func (q *Query) cleanQ() {
	if q.Q != "" {
		q.Answer, q.Decimal = parser.ReadNumber(q.Q[1:qNumberRange+1], q.Q[qNumberRange+2:28])
	}
}

func (q *Query) clean() {
	if q.IsValid() {
		q.I = strings.TrimPrefix(q.I, "I-")
		q.U = strings.TrimPrefix(q.U, "U-")
		q.M = strings.TrimPrefix(q.M, "M-")
		q.S = strings.TrimPrefix(q.S, "S-")
		q.Q = strings.TrimPrefix(q.Q, "Q-")
		q.E = strings.TrimPrefix(q.E, "E-")
		q.T = strings.TrimPrefix(q.T, "T-")
		q.cleanQ()
	}
}

func (q Query) IsValid() bool {
	return q.I != "" &&
		q.U != "" &&
		q.M != "" &&
		q.S != "" &&
		((q.Q != "" && q.E != "") || q.T != "")
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
