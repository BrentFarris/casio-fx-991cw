package decoder

import (
	"errors"
	"strings"
)

const (
	argCode    = "1A"
	endArgCode = "1B"
)

func CasioFX991CW() Decoder {
	d := New()
	d.AddEncoding(Encoding{Key: endArgCode, Transform: ")"})
	d.AddEncoding(Encoding{Key: "2E", Transform: "."})
	d.AddEncoding(Encoding{Key: "C0", Transform: "-"})
	d.AddEncoding(Encoding{Key: "C9" + argCode, Transform: "^("})
	d.AddEncoding(Encoding{Key: "A6", Transform: "+"})
	d.AddEncoding(Encoding{Key: "A7", Transform: "-"})
	d.AddEncoding(Encoding{Key: "A8", Transform: "×"})
	d.AddEncoding(Encoding{Key: "A9", Transform: "÷"})
	d.AddEncoding(Encoding{Key: "60", Transform: "("})
	d.AddEncoding(Encoding{Key: "D0", Transform: ")"})
	d.AddEncoding(Encoding{Key: "77", Transform: "sin("})
	d.AddEncoding(Encoding{Key: "78", Transform: "cos("})
	d.AddEncoding(Encoding{Key: "79", Transform: "tan("})
	d.AddEncoding(Encoding{Key: "7A", Transform: "sin^-1("})
	d.AddEncoding(Encoding{Key: "7B", Transform: "cos^-1("})
	d.AddEncoding(Encoding{Key: "7C", Transform: "tan^-1("})
	d.AddEncoding(Encoding{Key: "7D", Transform: "log("})
	d.AddEncoding(Encoding{Key: "74" + argCode, Transform: "sqrt("})
	d.AddEncoding(Encoding{Key: "75", Transform: "ln("})
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
