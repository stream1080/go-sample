package encrypt

import "testing"

func TestEncryptPassword(t *testing.T) {

	cases := []struct {
		str string
	}{
		{"e10adc3949ba59abbe56e057f20f883e"},
		{"af58e5f46a90448fbdc455aef5f20fb12"},
		{"f58e5f46a90448fbdc455aef5f20fb1"},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if got, err := EncryptPassword(c.str); err != nil {
				t.Errorf("EncryptPassword(%v) = %v, err: %v", c.str, got, err)
			}
		})
	}
}

func TestValidPassword(t *testing.T) {

	cases := []struct {
		str string
	}{
		{"e10adc3949ba59abbe56e057f20f883e"},
		{"af58e5f46a90448fbdc455aef5f20fb12"},
		{"f58e5f46a90448fbdc455aef5f20fb1"},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ciphertext, err := EncryptPassword(c.str)
			if err != nil {
				t.Errorf("EncryptPassword(%v) = %v, err: %v", c.str, ciphertext, err)
			}
			if !ValidPassword(c.str, ciphertext) {
				t.Errorf("ValidPassword(%v) = failed with err: %v", c.str, err)
			}
		})
	}
}
