package plates

import (
	"testing"
)

func TestAll(t *testing.T) {
	for code, gen := range Generators {
		t.Run(code, func(t *testing.T) {
			plates := make(map[string]bool)
			for i := 0; i < 100; i++ {
				plate := gen()
				t.Logf("Generated plate: %s", plate.Number)
				if plates[plate.Number] {
					t.Fatalf("Duplicated Plate(%d) - %s", i, plate.Number)
				}
				plates[plate.Number] = true
			}
		})
	}
}
