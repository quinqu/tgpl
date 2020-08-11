package converter

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meters float64

type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (feet Feet) String() string     { return fmt.Sprintf("%g feet", feet) }
func (meters Meters) String() string { return fmt.Sprintf("%g meters", meters) }

func (pounds Pounds) String() string       { return fmt.Sprintf("%g pounds", pounds) }
func (kilograms Kilograms) String() string { return fmt.Sprintf("%g kilograms", kilograms) }
