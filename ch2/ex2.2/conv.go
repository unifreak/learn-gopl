// Conv can converts temperature, length or weight units.
package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

func (c Celsius) String() string 	{ return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (ft Feet) String() string  	{ return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string 		{ return fmt.Sprintf("%gm", m) }
func (p Pound) String() string 		{ return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string 	{ return fmt.Sprintf("%gkg", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenhheit temperatur to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FtToM converts Feet to Meter.
func FtToM(ft Feet) Meter { return Meter(ft * 0.3048) }

// MToFt converts Meter to Feet.
func MToFt(m Meter) Feet { return Feet(m / 0.3048) }

// LbToKg converts Pound to Kilogram.
func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KgToLb converts Kilogram to Pound.
func KgToLb(k Kilogram) Pound { return Pound(k / 0.45359237) }