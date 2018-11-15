package nano

import (
	"bytes"
	"testing"

	"github.com/discv5/go-nano/nano/internal/util"
)

func TestNanoAddress(t *testing.T) {
	s1 := "xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuohr3"
	s2 := "xrb_1111111111111111111111111111111111111111111111111111hifc8npp"

	address, err := ParseAddress(s1)
	if err != nil {
		t.Fatal(err)
	}

	if address.String() != s1 {
		t.Fatalf("addresses are not equal")
	}

	address, err = ParseAddress(s2)
	if err != nil {
		t.Fatal(err)
	}

	zeroKey := util.MustDecodeHex("0000000000000000000000000000000000000000000000000000000000000000")
	if !bytes.Equal(address[:], zeroKey) {
		t.Fatalf("address is not zero")
	}
}
