package main

func dnaAlleles(s string) (string, error) {
	output := ""
	for _, letter := range s {
		swappedAllele, err := swapAllele(string(letter))
		if err != nil {
			return "", err
		}
		output = output + swappedAllele
	}
	return output, nil
}
