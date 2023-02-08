package conv

import "math"

// Konverterer Farhenheit til Celsius

func FarhenheitToCelsius(value float64) float64 {
		result := (value - 32) * (5.0 / 9.0)
		return math.Round(result*100) / 100
	}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64{
	result:= (value - 32.0) * (5.0/9.0) + 273.15
	return math.Round(result*100) / 100
}
// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64{
	result := value * (9.0/5.0) + 32.0
	return math.Round(result*100) / 100
}
// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64{
	result := value + 273.15
	return math.Round(result*100) / 100
}
// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64{
	result := (value - 273.15)*(9.0/5.0)+32.0
	return math.Round(result*100) / 100
}
// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64{
	result := value - 273.15
	return math.Round(result*100) / 100
}

