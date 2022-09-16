package plates

import (
	"math/rand"

	wr "github.com/mroth/weightedrand"
)

var DE_REGION_CHOOSER, _ = wr.NewChooser(
	wr.Choice[int]{Item: 1, Weight: 60},
	wr.Choice[int]{Item: 2, Weight: 35},
	wr.Choice[int]{Item: 3, Weight: 5},
)

func DE() Plate {
	l := rand.Intn(3) + 6
	number := String(DE_REGION_CHOOSER.Pick(), CHARS)
	number += " "
	number += String(l-len(number), ALPHANUMERICAL)

	return Plate{
		Country: "D",
		Number:  number,
	}
}
