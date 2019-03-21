package main

import "errors"

func swapAllele(letter string) (string, error) {
	matchedPairs := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	if _, ok := matchedPairs[letter]; ok {
		return matchedPairs[letter], nil
	}

	return "", errors.New("Not a valid Allele")
}
