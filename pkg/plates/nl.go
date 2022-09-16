package plates

import "math/rand"

var NL_REGIONS = []string{
	"A", "B", "D", "E", "G", "GZ", "GX", "H", "HZ", "HX", "K", "M", "N", "L", "P",
}

func NL() Plate {
	a, b := NUMERICAL, CHARS
	if rand.Intn(10)%3 == 0 {
		a, b = b, a
	}
	return Plate{
		Country: "NL",
		Number:  NL_REGIONS[rand.Intn(len(NL_REGIONS))] + "-" + String(3, a) + "-" + String(3, b),
	}
}
