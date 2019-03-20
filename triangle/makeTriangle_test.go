package main

import "testing"

func TestMakeTriangleReturnsTheCorrectlyPaddedStringDependingOnTheTotalNumberOfLines(t *testing.T) {
	result := makeTriangle(1)
	if len(result[0]) != 1 {
		t.Error("the length of the first line should be 1 when we are just building a single lined triangle")
	}
}

func TestTheWidthOfEachLineIsTheTotalNumberOfLinesDoubledMinusOne(t *testing.T) {
	result := makeTriangle(5)
	if len(result[4]) != 9 {
		t.Error("the final string wasn't passed correctly.\nExpected\t*****\nGot\t\t\t" + result[4])
	}
}
