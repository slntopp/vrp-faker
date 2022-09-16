package plates

func B() Plate {
	return Plate{
		"B",
		"1-" + String(3, CHARS) + "-" + String(3, NUMERICAL),
	}
}
