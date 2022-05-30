package conv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string 		{ return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string 	{ return fmt.Sprintf("%gkg", k) }

// LbToKg converts Pound to Kilogram.
func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KgToLb converts Kilogram to Pound.
func KgToLb(k Kilogram) Pound { return Pound(k / 0.45359237) }