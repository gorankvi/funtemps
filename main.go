package main

import (
	"flag"
	"fmt"

	"github.com/gorankvi/funtemps/conv"
)

// definisjon av variablene
var fahr float64
var celsius float64
var kelvin float64
var out string
var funfacts string
var temperatur string

// De forsjellige flaggdefinisjonen
func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "Kalkulerer temperatur i C-celsius, F-fahrenheit eller K-kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&temperatur, "t", "0", "temperaturskala C, K eller F for funfacts")
}

func main() {
	flag.Parse()
	// Gir feilmelding dersom noen prøver å legge inn flere flag for å konverteres.
	if (isFlagPassed("F") && isFlagPassed("C")) || (isFlagPassed("F") && isFlagPassed("K")) || (isFlagPassed("C") && isFlagPassed("K")) {
		fmt.Println("Error: -F, -C og -K kan ikke brukes samtidig.")
		return
	}

	// Konverterer fra celsius til kelvin, og kelvin til celsius
	if isFlagPassed("out") {
		if isFlagPassed("C") {
			if out == "K" {
				fmt.Printf("%.2f°C er %.2fK\n", celsius, conv.CelsiusToKelvin(celsius))
			}
		} else if isFlagPassed("K") {
			if out == "C" {
				fmt.Printf("%.2fK er %.2f°C\n", kelvin, conv.KelvinToCelsius(kelvin))
			}
		}
	}
	// Konverterer fra celsius til fahrenheit, og fra fahrenheit til celsius
	if isFlagPassed("out") {
		if isFlagPassed("C") {
			if out == "F" {
				fmt.Printf("%.2f°C er %.2fF°\n", celsius, conv.CelsiusToFahrenheit(celsius))
			}
		} else if isFlagPassed("F") {
			if out == "C" {
				fmt.Printf("%.2f°F er %.2fC°\n", fahr, conv.FahrenheitToCelsius(fahr))
			}
		}
	}
	// Konverterer fra fahrenheit til kelvin , og fra kelvin til fahrenheit
	if isFlagPassed("out") {
		if isFlagPassed("F") {
			if out == "K" {
				fmt.Printf("%.2f°F er %.2fK\n", fahr, conv.FahrenheitToKelvin(fahr))
			}
		} else if isFlagPassed("K") {
			if out == "F" {
				fmt.Printf("%.2fK er %.2fF°\n", kelvin, conv.KelvinToCelsius(fahr))
			}
		}
	}

}

func isFlagPassed(flagName string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			found = true
		}
	})
	return found
}
