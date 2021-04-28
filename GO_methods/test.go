package main

import "fmt"

type celsius float64 // Новый тип celsius
type kelvin float64  // Новый тип kelvin
type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5.0 / 9.0)
}
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {

	//var c celsius = 36.6
	var k kelvin = 500
	var f fahrenheit = 500
	fmt.Printf("%.3f форенгейт в цельсий\n  ", f.celsius())
	fmt.Printf("%.3f кельвин в цельсий\n  ", k.celsius())
}
