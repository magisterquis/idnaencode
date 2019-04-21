// Package idnaencoding encodes bytes as DNS IDNA labels.
//
// The encoded string will contain only letters, digits, and hyphens and is
// Suitable for use as a DNS label.
//
// Encoding works by adding a constant rune to each input byte and taking the
// resulting rune slice as a string to encode with punycode.
//
// Decoding is the reverse of Encoding
package idnaencoding

/*
 * idnaencoding.go
 * Encode bytes to IDNA labes
 * By J. Stuart McMurray
 * Created 20190421
 * Last Modified 20190421
 */

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

// DefaultOffset is the offset for the default Encoder
const DefaultOffset = 0xFF

// DefaultEncoder is an Encoder using DefaultOffset
var DefaultEncoder = Encoder(DefaultOffset)

// Encode wraps DefaultEncoder.Encode
func Encode(p []byte) (string, error) {
	return DefaultEncoder.Encode(p)
}

// Decode wraps DefaultEncoder.Decode
func Decode(s string) ([]byte, error) {
	return DefaultEncoder.Decode(s)
}

// An encoder is the value by which data is shifted to encode and decode it.
// It is not safe to change an Encoder's concurrently with calls to its
// Encode and Decode methods.
type Encoder rune

// Encode shifts p by the value of e and Punycodes it.
func (e Encoder) Encode(p []byte) (string, error) {
	/* Shift input */
	rs := make([]rune, len(p))
	for i, v := range p {
		rs[i] = rune(v) + rune(e)
	}

	/* Punycode */
	return idna.ToASCII(string(rs))
}

// Decode decodes a string encoded with Encode
func (e Encoder) Decode(s string) ([]byte, error) {
	/* Un-Punycode */
	u, err := idna.ToUnicode(s)
	if nil != err {
		return nil, err
	}

	/* Shift */
	var (
		d    rune
		next int
		b    = make([]byte, utf8.RuneCountInString(u))
	)
	for _, r := range u {
		/* Make sure we don't have a rune that's too big or too
		small */
		if r < rune(e) {
			return nil, fmt.Errorf(
				"decoded rune %+q below encoding bounds",
				r,
			)
		}
		d = r - rune(e)
		if 0xff < d {
			return nil, fmt.Errorf("decoded rune %+q too large", r)
		}
		b[next] = byte(d)
		next++
	}

	return b, nil
}
