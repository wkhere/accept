package accept

import "testing"

var aeTab = []struct {
	header, val string
	want        bool
}{
	{"", "", false},
	{"", "any", false},
	{"foo", "", false},
	{"foo, bar", "", false},
	{"foo, bar", "FOO", true},
	{"FOO, BAR", "foo", true},
	{"FOO, BAR", "bar", true},
	{"x", "x", true},
	{"x", "y", false},
	{"x y", "x", false},
	{"x,y", "x", true},
	{"x,y", "y", true},
	{"x,y,z", "y", true},
	{"x, y", "x", true},
	{"x, y", "y", true},
	{"x,y,z", "y", true},
	{"x, y, z", "y", true},
	{"x,y , z", "y", true},
	{"x, y , z", "y", true},
	{"x,y,z", "bad", false},
	{"x;q=0.4", "x", true},
	{"x;q=0.4,y", "x", true},
	{"x;q=0.4, y", "x", true},
	{"x; q=0.4", "x", true},
	{"x; q=0.4,", "x", true},
	{"x; q=0.4 ,", "x", true},
	{"x; q= 0.4", "x", false},
	{"x; q =0.4", "x", false},
	{"x; q=0.1err ,", "x", false},
	{"x; q=err ,", "x", false},
	{"x; q", "x", false},
	{"x; q= ,", "x", false},
	{"x; q=1", "x", true},
	{"x; q=1", "x", true},
	{"x; q=0", "x", false},
	{"x; foo", "x", false},
	{"x; foo=2", "x", false},
	{"x; q=0.0", "x", false},
	{"x; q=0,y", "x", false},
	{"x; q=0, y", "x", false},
	{"x; q=err, y", "y", true},
	{"x; q=, y", "y", true},
	{"x; q, y", "y", true},
	{"x; q=0, y", "y", true},
	{"x; q=0.0, y", "y", true},
	{"*", "x", true},
	{"*, redundant", "x", true},
	{"*; q=1", "x", true},
	{"*; q=0.5", "x", true},
	{"*; q=0", "x", false},
	{"*; q=0.0", "x", false},
	{"*; q=", "x", false},
	{"*; q=0.1err", "x", false},
	{"*; q=1err", "x", false},
	{"*; q=err", "x", false},
	{"*", "", false},
	// not testing "identity" as I don't really undestand some cases there
}

func TestAcceptEncoding(t *testing.T) {
	for i, tc := range aeTab {
		have := Encoding(tc.header, tc.val)
		if have != tc.want {
			t.Errorf("tc#%d: have %v, want %v", i, have, tc.want)
		}
	}
}

func BenchmarkAcceptEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range aeTab[:50] {
			_ = Encoding(tc.header, tc.val)
		}
	}
}
