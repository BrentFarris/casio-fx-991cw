package decoder

import (
	"errors"
	"strings"
)

const (
	argCode    = "1A"
	endArgCode = "1B"
)

type CasioFX991CWTable struct {
	Columns [5][]string
}

func CasioFX991CW() Decoder {
	d := New()
	d.AddEncoding(Encoding{Key: endArgCode, Transform: ")"})
	d.AddEncoding(Encoding{Key: "21", Transform: "e"})
	d.AddEncoding(Encoding{Key: "22", Transform: "π"})
	d.AddEncoding(Encoding{Key: "2E", Transform: "."})
	d.AddEncoding(Encoding{Key: "A6", Transform: "+"})
	d.AddEncoding(Encoding{Key: "A7", Transform: "-"})
	d.AddEncoding(Encoding{Key: "A8", Transform: "×"})
	d.AddEncoding(Encoding{Key: "A9", Transform: "÷"})
	d.AddEncoding(Encoding{Key: "C0", Transform: "-"})
	d.AddEncoding(Encoding{Key: "DD", Transform: "E"})
	d.AddEncoding(Encoding{Key: "DE", Transform: "P"})
	d.AddEncoding(Encoding{Key: "DF", Transform: "T"})
	d.AddEncoding(Encoding{Key: "E0", Transform: "G"})
	d.AddEncoding(Encoding{Key: "E1", Transform: "M"})
	d.AddEncoding(Encoding{Key: "E2", Transform: "k"})
	d.AddEncoding(Encoding{Key: "E3", Transform: "m"})
	d.AddEncoding(Encoding{Key: "E4", Transform: "µ"})
	d.AddEncoding(Encoding{Key: "E5", Transform: "n"})
	d.AddEncoding(Encoding{Key: "E6", Transform: "p"})
	d.AddEncoding(Encoding{Key: "E7", Transform: "f"})
	d.AddEncoding(Encoding{Key: "60", Transform: "("})
	d.AddEncoding(Encoding{Key: "D0", Transform: ")"})
	d.AddEncoding(Encoding{Key: "C9" + argCode, Transform: "^("})
	d.AddEncoding(Encoding{Key: "6C", Transform: "sinh("})
	d.AddEncoding(Encoding{Key: "6D", Transform: "cosh("})
	d.AddEncoding(Encoding{Key: "6E", Transform: "tanh("})
	d.AddEncoding(Encoding{Key: "6F", Transform: "sinh^-1("})
	d.AddEncoding(Encoding{Key: "70", Transform: "cosh^-1("})
	d.AddEncoding(Encoding{Key: "71", Transform: "tanh^-1("})
	d.AddEncoding(Encoding{Key: "77", Transform: "sin("})
	d.AddEncoding(Encoding{Key: "78", Transform: "cos("})
	d.AddEncoding(Encoding{Key: "79", Transform: "tan("})
	d.AddEncoding(Encoding{Key: "7A", Transform: "sin^-1("})
	d.AddEncoding(Encoding{Key: "7B", Transform: "cos^-1("})
	d.AddEncoding(Encoding{Key: "7C", Transform: "tan^-1("})
	d.AddEncoding(Encoding{Key: "7D", Transform: "log("})
	d.AddEncoding(Encoding{Key: "74" + argCode, Transform: "√("})
	d.AddEncoding(Encoding{Key: "75", Transform: "ln("})
	d.AddEncoding(Encoding{Key: "FD35", Transform: "h"})
	d.AddEncoding(Encoding{Key: "FD38", Transform: "ћ"})
	d.AddEncoding(Encoding{Key: "FD48", Transform: "c"})
	d.AddEncoding(Encoding{Key: "FD4F", Transform: "ε0"})
	d.AddEncoding(Encoding{Key: "FD50", Transform: "μ0"})
	d.AddEncoding(Encoding{Key: "FD54", Transform: "Zo"})
	d.AddEncoding(Encoding{Key: "FD56", Transform: "G"})
	d.AddEncoding(Encoding{Key: "FD58", Transform: "lP"})
	d.AddEncoding(Encoding{Key: "FD59", Transform: "tP"})
	d.AddEncoding(Encoding{Key: "FD36", Transform: "µN"})
	d.AddEncoding(Encoding{Key: "FD37", Transform: "µB"})
	d.AddEncoding(Encoding{Key: "FD46", Transform: "e"})
	d.AddEncoding(Encoding{Key: "FD51", Transform: "Φ0"})
	d.AddEncoding(Encoding{Key: "FD53", Transform: "G0"})
	d.AddEncoding(Encoding{Key: "FD5A", Transform: "KJ"})
	d.AddEncoding(Encoding{Key: "FD5B", Transform: "RK"})
	d.AddEncoding(Encoding{Key: "FD30", Transform: "mP"})
	d.AddEncoding(Encoding{Key: "FD31", Transform: "mn"})
	d.AddEncoding(Encoding{Key: "FD32", Transform: "me"})
	d.AddEncoding(Encoding{Key: "FD33", Transform: "mµ"})
	d.AddEncoding(Encoding{Key: "FD34", Transform: "ao"})
	d.AddEncoding(Encoding{Key: "FD39", Transform: "α"})
	d.AddEncoding(Encoding{Key: "FD3A", Transform: "re"})
	d.AddEncoding(Encoding{Key: "FD3B", Transform: "λc"})
	d.AddEncoding(Encoding{Key: "FD3C", Transform: "γP"})
	d.AddEncoding(Encoding{Key: "FD3D", Transform: "λcP"})
	d.AddEncoding(Encoding{Key: "FD3E", Transform: "λcn"})
	d.AddEncoding(Encoding{Key: "FD3F", Transform: "R∞"})
	d.AddEncoding(Encoding{Key: "FD41", Transform: "µP"})
	d.AddEncoding(Encoding{Key: "FD42", Transform: "µe"})
	d.AddEncoding(Encoding{Key: "FD43", Transform: "µn"})
	d.AddEncoding(Encoding{Key: "FD44", Transform: "µµ"})
	d.AddEncoding(Encoding{Key: "FD5C", Transform: "mT"})
	d.AddEncoding(Encoding{Key: "FD40", Transform: "mu"})
	d.AddEncoding(Encoding{Key: "FD45", Transform: "F"})
	d.AddEncoding(Encoding{Key: "FD47", Transform: "NA"})
	d.AddEncoding(Encoding{Key: "FD48", Transform: "K"})
	d.AddEncoding(Encoding{Key: "FD49", Transform: "Vm"})
	d.AddEncoding(Encoding{Key: "FD4A", Transform: "R"})
	d.AddEncoding(Encoding{Key: "FD4C", Transform: "C1"})
	d.AddEncoding(Encoding{Key: "FD4D", Transform: "C2"})
	d.AddEncoding(Encoding{Key: "FD4E", Transform: "σ"})
	d.AddEncoding(Encoding{Key: "FD52", Transform: "gn"})
	d.AddEncoding(Encoding{Key: "FD57", Transform: "atm"})
	d.AddEncoding(Encoding{Key: "FD5D", Transform: "RK-g0"})
	d.AddEncoding(Encoding{Key: "FD5E", Transform: "Kj-g0"})
	d.AddEncoding(Encoding{Key: "FD55", Transform: "t"})
	d.AddEncoding(Encoding{
		Key: "7D" + argCode,
		TransformFunc: func(in string, out *strings.Builder) (int, error) {
			out.WriteString("log")
			read := 0
			for len(in) >= 2 && !strings.HasPrefix(in, "1C") {
				if in[0] != '3' {
					return 0, errors.New("currently only support constant log base")
				}
				out.WriteRune([]rune(in)[1])
				in = in[2:]
				read += 2
			}
			read += len("1C")
			out.WriteString("(")
			return read, nil
		},
	})
	d.AddEncoding(Encoding{
		Key: "3",
		TransformFunc: func(in string, out *strings.Builder) (int, error) {
			if in == "" {
				return 0, errors.New("expected a constant after '3' encoding")
			}
			out.WriteRune([]rune(in)[0])
			return 1, nil
		},
	})
	return d
}
