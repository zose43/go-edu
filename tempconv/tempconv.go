package tempconv

// Package tempConv performs Fahrenheits and Celsius conversions

import "fmt"

type Celsius float32
type Fahrenheit float32
type Kelvin float32

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

func (c Celsius) String() string {
	return fmt.Sprintf("%f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%f℉", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%fK", k)
}
