package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func FToK(f Fahrenheit) Kelvin {
	c := (FToC(f))
	return Kelvin(c - AbsoluteZeroC)
}

func KToF(k Kelvin) Fahrenheit {
	c := KToC(k)
	return Fahrenheit(CToF(c))
}

func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }
