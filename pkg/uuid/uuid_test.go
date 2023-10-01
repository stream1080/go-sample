package uuid

import "testing"

func TestNew(t *testing.T) {
	uuid := New()
	if len(uuid) <= 0 {
		t.Fatal("new uuid failed")
	}
	t.Log(uuid)
}

func TestUUID16(t *testing.T) {
	uuid := UUID16()
	if len(uuid) <= 0 {
		t.Fatal("new uuid 16 failed")
	}
	t.Log(uuid)
}
