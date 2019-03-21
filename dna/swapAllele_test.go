package main

import "testing"

func TestOnlyValidAllelesCanBeSwapped(t *testing.T) {
	_, err := swapAllele("Z")
	if err == nil {
		t.Errorf("Z is not a valid allele")
	}

	swapped, err := swapAllele("A")
	if err != nil {
		t.Errorf("A is not a valid allele")
	}
	if swapped != "T" {
		t.Errorf("Allele A was not correctly converted to T")
	}
}
