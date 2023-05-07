package tempconv

// Package tempConv performs Fahrenheits and Celsius conversions

import (
	"flag"
	"fmt"
)

type Celsius float32
type Fahrenheit float32
type Kelvin float32
type celsiusFlag struct {
	Celsius
}

const (
	AbsoluteZeroC Celsius    = -273.15
	AbsoluteZeroK Kelvin     = -273.15
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingC     Celsius    = 0
	FreezingK     Kelvin     = 273.2
	FreezingF     Fahrenheit = 32
	BoilingC      Celsius    = 100
	BoilingK      Kelvin     = 373.2
	BoilingF      Fahrenheit = 212
)

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "°C", "C":
		f.Celsius = Celsius(value)
		return nil
	case "°F", "F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("Invalid temp %q\n", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (c Celsius) String() string {
	return fmt.Sprintf("%f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%f℉", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%fK", k)
}
