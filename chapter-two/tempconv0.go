package main

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	c := FToC(212.0)
	fmt.Println(c)

}

func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
