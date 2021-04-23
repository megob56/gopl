package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Farenheit float64
type Kelvin float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v/n", err)
			os.Exit(1)
		}
		f := Farenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))

		ft := Feet(t)
		m := Meters(t)
		fmt.Printf("%s = %s, %s = %s\n", ft, FtToM(ft), m, MToFt(m))
	}
}

func CToF(c Celsius) Farenheit         { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celsius         { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsius            { return Celsius(k - 273.15) }
func KToF(k Kelvin) Farenheit          { return Farenheit((k-273.15)*9/5 + 32) }
func FtToM(ft Feet) Meters             { return Meters(ft * 0.3) }
func MToFt(m Meters) Feet              { return Feet(m * 3.2) }
func PoundsToKilos(p Pounds) Kilograms { return Kilograms(p / 2.2) }
func KilosToPounds(k Kilograms) Pounds { return Pounds(k * 2.2) }
