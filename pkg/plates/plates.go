package plates

import "math/rand"

type Plate struct {
	Country string
	Number  string
}

const CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CROSSING = "ABCEHIKMOPTXY" // Chars which are common between latin and cyrillic
const ALPHANUMERICAL = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const NUMERICAL = "0123456789"

type Generator func() Plate

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int, letters string) string {
	return StringWithCharset(length, letters)
}
