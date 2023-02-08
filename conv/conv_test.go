package conv

import (
	"reflect"
	"testing"
)

// Tester at funksjonen FarhenheitToCelsius gir 56.67 ved input 134: PASS
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Tester at FarhenheitToKelvin gir 329.82 ved input 134.01: PASS
func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134.01, want: 329.82},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
// Tester at CelsiusTOFahrenheit gir 134.01 ved input 56.67: PASS
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134.01},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
// Tester at CelsiusToKelvin gir 329.82 ved input 56.67: PASS
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 329.82},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
// Tester at KelvinToFarhenheit gir 134.01 ved input 329.82: PASS
func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 134.01},
	}

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
// Tester at KelvinToCelsius gir 56.67 ved input 329.82: PASS
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.67},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}