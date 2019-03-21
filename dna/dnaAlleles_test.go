package main

import "testing"

func TestSwappedDnaStringsWorksCorrectly(t *testing.T) {
	expectedStrings := map[string]string{
		"ATTCG": "TAAGC",
	}

	for input, expectedOutput := range expectedStrings {
		output, _ := dnaAlleles(input)
		if output != expectedOutput {
			t.Errorf("DNS sequence not converted correctly")
		}
	}

	_, err := dnaAlleles("NOT VALID")
	if err == nil {
		t.Errorf("'NOT VALID' is not a valid DNA strand, dummy")
	}
}
