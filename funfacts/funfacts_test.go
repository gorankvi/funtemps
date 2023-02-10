package funfacts

import (
	"reflect"
	"testing"
)
// Tester at GetFunFacts returnerer riktig string ved input. PASS
func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "sun", want: facts.Sun},
		{input: "luna", want: facts.Luna},
		{input: "terra", want: facts.Terra},
		{input: "unknown", want: []string{}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
