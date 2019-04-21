package idnaencoding

/*
 * idna_encoding_test.go
 * Tests for idnaencoding
 * By J. Stuart McMurray
 * Created 20190421
 * Last Modified 20190421
 */

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncoder(t *testing.T) {
	for _, c := range []struct {
		have []byte
		want string
	}{
		{[]byte{0x00}, "xn--wda"},
		{[]byte("abc123"), "xn--bfacd5pfg"},
		{make([]byte, 10), "xn--wdaaaaaaaaaa"},
	} {
		gotE, err := Encode(c.have)
		if nil != err {
			t.Fatalf("Encode: %v", err)
		}
		if c.want != gotE {
			t.Fatalf(
				"Encode: have:%02x got:%v want:%v",
				c.have,
				gotE,
				c.want,
			)
		}

		gotD, err := Decode(gotE)
		if nil != err {
			t.Fatalf("Decode: %v", err)
		}
		if 0 != bytes.Compare(c.have, gotD) {
			t.Fatalf(
				"Decode have:%v got:%02x want:%02x",
				gotE,
				gotD,
				c.have,
			)
		}
	}
}

func ExampleEncode() {
	/* Slice of arbitrary bytes */
	bs := []byte("These are abitrary bytes: \x13\xA3\x99\xF8")
	/* Encode it */
	s, err := Encode(bs)
	if nil != err {
		panic(err)
	}
	fmt.Printf("%s\n", s)

	// Output: xn--geazaaa9m2g1dbakc5aacg3b0a0ofbkprf1kc61gji54g
}
