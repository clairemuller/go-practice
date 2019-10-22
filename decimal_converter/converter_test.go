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

func TestFormatHex(t *testing.T) {
	tables := []struct {
		decimal float64
		hex     string
	}{
		{-9, "Whole numbers only!"},
		{7.4, "Whole numbers only!"},
		{-83, "Whole numbers only!"},
		{8623.5435, "Whole numbers only!"},
		{0, "0"},
		{7, "7"},
		{11, "B"},
		{18, "12"},
		{39, "27"},
		{63, "3F"},
		{99, "63"},
		{176, "B0"},
		{335, "14F"},
		{738, "2E2"},
		{4095, "FFF"},
		{66654, "1045E"},
		{91757843, "5781D13"},
	}

	for _, table := range tables {
		hexNum := FormatHex(table.decimal)
		if hexNum != table.hex {
			t.Errorf("FormatHex was incorrect, given: %v, got: %v, want: %v", table.decimal, hexNum, table.hex)
		}
	}
}
