package main

import (
	file "aoc23/internal"
	"slices"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	var hands []Hand
	for _, line := range input {
		hands = append(hands, NewHand(line, false))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return CompareHands(a, b)
	})
	total := Part1Solver(hands)
	var expected int = 6440
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestPartTwo(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	total := Part2Solver(input)
	var expected int = 0
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestGetType(t *testing.T) {
	tests := []struct {
		hand     []byte
		expected int
	}{
		{[]byte("AAAAA"), FIVE_OF_A_KIND},
		{[]byte("AA8AA"), FOUR_OF_A_KIND},
		{[]byte("23332"), FULL_HOUSE},
		{[]byte("TTT98"), THREE_OF_A_KIND},
		{[]byte("23432"), TWO_PAIR},
		{[]byte("A23A4"), ONE_PAIR},
		{[]byte("23456"), HIGH_CARD},
	}
	for _, test := range tests {
		result := GetType(test.hand, false)
		if result != test.expected {
			t.Errorf("Expected %v, but got %v for %v", test.expected, result, string(test.hand))
		}
	}
}

func TestDoesHand1Win(t *testing.T) {
	tests := []struct {
		hand1    Hand
		hand2    Hand
		expected int
	}{
		{NewHand([]byte("33332"), false), NewHand([]byte("2AAAA"), false), 1},
		{NewHand([]byte("77888"), false), NewHand([]byte("77788"), false), 1},
		{NewHand([]byte("2AAAA"), false), NewHand([]byte("33332"), false), -1},
		{NewHand([]byte("77788"), false), NewHand([]byte("77888"), false), -1},
	}
	for _, test := range tests {
		result := CompareHands(test.hand1, test.hand2)
		if result != test.expected {
			t.Errorf("Expected hand %v and hand %v to return %v", string(test.hand1.Cards), string(test.hand2.Cards), test.expected)
		}
	}
}