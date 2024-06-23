package decoder

import "strings"

func CasioFX991CW() Decoder {
	d := New()
	d.AddEncoding(Encoding{Key: "2E", Transform: "."})
	d.AddEncoding(Encoding{Key: "C91A", Transform: "^("})
	d.AddEncoding(Encoding{Key: "1B", Transform: ")"})
	d.AddEncoding(Encoding{Key: "A6", Transform: "+"})
	d.AddEncoding(Encoding{Key: "A7", Transform: "-"})
	d.AddEncoding(Encoding{Key: "A8", Transform: "×"})
	d.AddEncoding(Encoding{Key: "A9", Transform: "÷"})
	d.AddEncoding(Encoding{Key: "603", Transform: "("})
	d.AddEncoding(Encoding{Key: "D0", Transform: ")"})
	d.AddEncoding(Encoding{Key: "773", Transform: "sin("})
	d.AddEncoding(Encoding{Key: "783", Transform: "cos("})
	d.AddEncoding(Encoding{Key: "793", Transform: "tan("})
	d.AddEncoding(Encoding{Key: "7A", Transform: "sin^-1("})
	d.AddEncoding(Encoding{Key: "7B", Transform: "cos^-1("})
	d.AddEncoding(Encoding{Key: "7C", Transform: "tan^-1("})
	d.AddEncoding(Encoding{Key: "7D", Transform: "log("})
	d.AddEncoding(Encoding{Key: "7D1A", Transform: "log"})
	d.AddEncoding(Encoding{Key: "741A", Transform: "sqrt("})
	d.AddEncoding(Encoding{Key: "1C", Transform: ")"}) // 7D1A ending
	d.AddEncoding(Encoding{
		Key: "3",
		TransformFunc: func(in string, out *strings.Builder) int {
			out.WriteRune([]rune(in)[0])
			return 1
		},
	})
	return d
}
