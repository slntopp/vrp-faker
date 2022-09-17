package plates

import "math/rand"

var UA_REGIONS = []string{
	"KA", "KB", "KC", "KE", "KH", "KI", "KK", "KM", "KO", "KP", "KT", "KX", "HA", "HB", "HC", "HE", "HH", "HI", "HK", "HM", "HO", "HP", "HT", "HX", "IA", "IB", "IC", "IE", "IH",
}

func UA() Plate {
	return Plate{
		Country: "UA",
		Number:  UA_REGIONS[rand.Intn(len(UA_REGIONS))] + " " + String(4, NUMERICAL) + " " + String(2, CROSSING),
	}
}
