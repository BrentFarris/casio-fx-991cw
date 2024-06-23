package decoder

import (
	"casiofx991cw/query"
	"testing"
)

type Test struct {
	Expected string
	URL      string
}

func TestCasioFX991CW(t *testing.T) {
	tests := []Test{
		{
			Expected: "3×9=27",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-02700000000000000000000001010000000000000000000000000000+E-33A839",
		},
		{
			Expected: "4×4=16",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01600000000000000000000001010000000000000000000000000000+E-34A834",
		},
		{
			Expected: "2×5=10",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01000000000000000000000001010000000000000000000000000000+E-32A835",
		},
		{
			Expected: "7×6=42",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-04200000000000000000000001010000000000000000000000000000+E-37A836",
		},
		{
			Expected: "8×8=64",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-06400000000000000000000001010000000000000000000000000000+E-38A838",
		},
		{
			Expected: "9=9",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000001000000000000000000000000000000+E-39",
		},
		{
			Expected: "9×9=81",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-08100000000000000000000001010000000000000000000000000000+E-39A839",
		},
		{
			Expected: "3+7=10",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01000000000000000000000001010000000000000000000000000000+E-33A637",
		},
		{
			Expected: "319+1=320",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-03200000000000000000000001020000000000000000000000000000+E-333139A631",
		},
		{
			Expected: "3200=3200",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-03200000000000000000000001030000000000000000000000000000+E-33323030",
		},
		{
			Expected: "0.0009=0.0009", // 9×10^-4
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000000960000000000000000000000000000+E-302E30303039",
		},
		{
			Expected: "0.09=0.09",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-09000000000000000000000000980000000000000000000000000000+E-302E3039",
		},
		{
			Expected: "3÷4=0.75",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-07500000000000000000000000990000000000000000000000000000+E-33A934",
		},
		{
			//"123456789123456789=1.234567891x10^17",
			Expected: "123456789123456789=123456789123456789",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01234567891234567890000001170000000000000000000000000000+E-313233343536373839313233343536373839",
		},
		{
			Expected: "35÷555=0.063063063063063063063063",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-06306306306306306306306300980000000000000000000000000000+E-3335A9353535",
		},
		{
			Expected: "2^(2)=4",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-04000000000000000000000001000000000000000000000000000000+E-32C91A321B",
		},
		{
			Expected: "sqrt(4)=2",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-02000000000000000000000001000000000000000000000000000000+E-741A341B",
		},
		{
			Expected: "3+6(5+4)=57",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-05700000000000000000000001010000000000000000000000000000+E-33A6366035A634D0",
		},
		{
			Expected: "log2(32)=5",
			URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-05000000000000000000000001000000000000000000000000000000+E-7D1A321C33321B",
		},
		/*
			{
				Expected: "1+2-3×4÷(5(6))+sin(1)+cos(2)+tan(3)=2.882777605",
				URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-02882777605186476314359301000000000000000000000000000000+E-31A632A733A834A960356036D0D0A67731D0A67832D0A67933D0",
			},
			{
				Expected: "ln(5)=1.609437912",
				URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-01609437912434100374600701000000000000000000000000000000+E-7535D0",
			},
			{
				Expected: "5(3/2)=6.5",
				URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-26A1A200000000000000000001050000000000000000000000000000+E-181F1D1A351B1A331B1A321B1E",
			},
			{
				Expected: "3/4=0.75",
				URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-23A40000000000000000000001030000000000000000000000000000+E-C81D1A331B1A341B1E",
			},
			{
				Expected: "123/456=0.2697368421",
				URL:      "http://wes.casio.com/ncal/index.php?q=I-005A+U-801100018E66+M-C10000AA00+S-001510100010001A1010B00006CC+Q-241A1520000000000000000001060000000000000000000000000000+E-C81D1A3132331B1A3435361B1E",
			},
		*/
	}
	d := CasioFX991CW()
	for i := range tests {
		q := query.New(tests[i].URL)
		if !q.IsValid() {
			t.Errorf("test %d '%s' failed due to invalid query", i, tests[i].Expected)
		}
		res, err := d.Decode(q)
		if err != nil {
			t.Error(err.Error() + ". Current: " + res)
		} else if res != tests[i].Expected {
			t.Errorf("test %d failed; expected '%s' but got '%s'", i, tests[i].Expected, res)
		}
	}
}
