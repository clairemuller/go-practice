package main

import "testing"

func TestFormatBinary(t *testing.T) {
	tables := []struct {
		decimal float64
		binary  string
	}{
		{-9, "Whole numbers only!"},
		{7.4, "Whole numbers only!"},
		{-83, "Whole numbers only!"},
		{8623.5435, "Whole numbers only!"},
		{0, "0"},
		{7, "111"},
		{18, "10010"},
		{39, "100111"},
		{63, "111111"},
		{99, "1100011"},
		{335, "101001111"},
	}

	for _, table := range tables {
		binaryNum := FormatBinary(table.decimal)
		if binaryNum != table.binary {
			t.Errorf("FormatBinary was incorrect, given: %v, got: %v, want: %v", table.decimal, binaryNum, table.binary)
		}
	}
}
