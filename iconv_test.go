package iconv

import "testing"

var tstData1 = []struct {
	srcChrTyp, dstChrTyp string
	src, out             string
}{
	{
		srcChrTyp: "UTF8", dstChrTyp: "utf-8",
		src: "1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000",
		out: "1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000",
	},
	{
		srcChrTyp: "UTF8", dstChrTyp: "utf-8",
		src: "1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000",
		out: "1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000" +
			"1111111111111111111111111111111111111111111111111111111111111111" +
			"2222222222222222222222222222222222222222222222222222222222222222" +
			"3333333333333333333333333333333333333333333333333333333333333333" +
			"4444444444444444444444444444444444444444444444444444444444444444" +
			"5555555555555555555555555555555555555555555555555555555555555555" +
			"6666666666666666666666666666666666666666666666666666666666666666" +
			"7777777777777777777777777777777777777777777777777777777777777777" +
			"8888888888888888888888888888888888888888888888888888888888888888" +
			"9999999999999999999999999999999999999999999999999999999999999999" +
			"0000000000000000000000000000000000000000000000000000000000000000",
	},
	{
		srcChrTyp: "UTF8", dstChrTyp: "ISO-8859-1",
		src: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" +
			"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC" +
			"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd" +
			"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" +
			"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF" +
			"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG" +
			"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
			"iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii" +
			"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" +
			"kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk" +
			"llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll" +
			"mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm" +
			"nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn" +
			"oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo" +
			"pppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppp" +
			"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		out: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" +
			"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC" +
			"dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd" +
			"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" +
			"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF" +
			"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG" +
			"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" +
			"iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii" +
			"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" +
			"kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk" +
			"llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll" +
			"mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm" +
			"nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn" +
			"oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo" +
			"pppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppp" +
			"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
	},
}

func TestIconv1(t *testing.T) {
	var n int
	var b []byte
	var outbuf [512]byte
	for i, d := range tstData1 {
		cd, err := Open(d.srcChrTyp, d.dstChrTyp)
		b, n, err = cd.Conv([]byte(d.src), outbuf[:])
		if err != nil {
			t.Fatalf("Test iconv1 return error non nil: %v", err)
		}
		if d.out != string(b) {
			t.Errorf("Wanted:%s\nGot:%s\nTest-Iconv1 test #%d failed with n=%d",
				d.out, string(b), i+1, n)
		}
	}
}