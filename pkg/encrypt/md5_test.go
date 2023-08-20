package encrypt

import (
	"testing"
)

func TestMd5Salt(t *testing.T) {

	cases := []struct {
		str      string
		salt     string
		expected string
	}{
		{"e10adc3949ba59abbe56e057f20f883e", "123456", "0871bdcd552962c1e3d393756ecec659"},
		{"af58e5f46a90448fbdc455aef5f20fb12", "dc89b1", "615d1b04158590ef74a2dab45bb48733"},
		{"f58e5f46a90448fbdc455aef5f20fb1", "work", "e4c99a7d3a594050da51e40ae48ac16c"},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if got := Md5Salt(c.str, c.salt); got != c.expected {
				t.Errorf("Md5Salt(%v, %v) = %v, want %v", c.str, c.salt, got, c.expected)
			}
		})
	}
}

func TestMd5(t *testing.T) {

	cases := []struct {
		str      string
		expected string
	}{
		{"e10adc3949ba59abbe56e057f20f883e", "14e1b600b1fd579f47433b88e8d85291"},
		{"af58e5f46a90448fbdc455aef5f20fb12", "659fcf634af3738a14b28b0efa58ddc6"},
		{"f58e5f46a90448fbdc455aef5f20fb1", "10cf4a9bc38d4d055b94c13c9e915a3b"},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if got := Md5(c.str); got != c.expected {
				t.Errorf("Md5(%v) = %v, want %v", c.str, got, c.expected)
			}
		})
	}
}
