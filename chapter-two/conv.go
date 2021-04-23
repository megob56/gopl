package tempconv

func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsius    { return Celsius(k - 273.15) }
func KToF(k Kelvin) Farenheit  { return Farenheit((k-273.15)*9/5 + 32) }
