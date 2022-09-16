package plates

func BY() Plate {
	return Plate{
		Country: "BY",
		Number:  String(4, NUMERICAL) + " " + String(2, CHARS) + "-" + String(1, "1234567"),
	}
}
