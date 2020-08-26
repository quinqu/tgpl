package converter

// celcius + fahreheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// feet + meters
func FeetToMeters(feet Feet) Meters { return Meters(feet / 3.281) }

func MetersToFeet(meters Meters) Feet { return Feet(meters * 3.281) }

// pounds + kilograms
func PoundsToKilograms(pounds Pounds) Kilograms { return Kilograms(pounds / 2.205) }

func KilogramsToPounds(kilograms Kilograms) Pounds { return Pounds(kilograms * 2.205) }
