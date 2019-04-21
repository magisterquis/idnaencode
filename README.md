IDNA Encoding
=============
Little library to encode arbitary byte slices to strings containing only
letters, digits, and hyphens.

Each byte of the string to be encoded is taken as a unicode code point and
shifted up by a set value.  The resultant code points are taken as a unicode
string and encoded using [Punycode](https://en.wikipedia.org/wiki/Punycode).

Example
-------
```go
func SendHardwareAddresses(domain string) {
	is, err := net.Interfaces()
	if nil != err {
		panic(err)
	}
	for _, i := range is {
		s, err := idnaencoding.Encode(i.HardwareAddr)
		if nil != err {
			panic(err)
		}
		if 0 != len(i.HardwareAddr) {
			go net.LookupHost(s + "." + domain)
		}
	}
}
```
