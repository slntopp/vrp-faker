package vrp

import p "github.com/slntopp/vrp-faker/pkg/plates"

var Generators = map[string]p.Generator{
	"B":  p.B,
	"BY": p.BY,
	"DE": p.DE,
	"NL": p.NL,
	"UA": p.UA,
}
