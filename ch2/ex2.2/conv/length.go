package conv

import "fmt"

type Feet float64
type Meter float64

func (ft Feet) String() string  	{ return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string 		{ return fmt.Sprintf("%gm", m) }

// FtToM converts Feet to Meter.
func FtToM(ft Feet) Meter { return Meter(ft * 0.3048) }

// MToFt converts Meter to Feet.
func MToFt(m Meter) Feet { return Feet(m / 0.3048) }
