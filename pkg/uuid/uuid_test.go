package uuid

import "testing"

func TestNew(t *testing.T) {
	uuid := New()
	if len(uuid) <= 0 {
		t.Fatal("new uuid failed")
	}
	t.Log(uuid)
}
