// +build linux

package netlink

import (
	"testing"
)

func TestFouAddDel(t *testing.T) {
	skipUnlessRoot(t)

	fou1 := NewFou()
	fou1.Port = 2222
	fou1.Gue = true
	if err := FouAdd(fou1); err != nil {
		t.Fatal(err)
	}

	if err := FouAdd(fou1); err == nil {
		t.Fatal("Could add fou with same port twice")
	}

	fou2 := NewFou()
	fou2.Port = 3333
	fou2.IpProto = 47

	if err := FouAdd(fou2); err != nil {
		t.Fatal(err)
	}

	if err := FouDel(fou1); err != nil {
		t.Fatal(err)
	}

	if err := FouDel(fou2); err != nil {
		t.Fatal(err)
	}
}