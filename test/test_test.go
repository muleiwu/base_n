package test

import (
	"testing"

	"github.com/muleiwu/base_n"
)

func TestBase62(t *testing.T) {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newBaseN := base_n.NewBaseN([]byte(str))

	base_n.NewBase30oO()
	encode := newBaseN.Encode(62)
	decode, err := newBaseN.Decode(encode)

	if err != nil {
		t.Fatal(err)
	}

	if decode != 62 {
		t.Errorf("Expected decode result to be 62, got %d", decode)
	}
}

func TestBase30oO(t *testing.T) {
	str := "0oO"
	newBaseN := base_n.NewBaseN([]byte(str))

	encode := newBaseN.Encode(5)
	decode, err := newBaseN.Decode(encode)

	if err != nil {
		t.Fatal(err)
	}

	if decode != 5 {
		t.Errorf("Expected decode result to be 5, got %d", decode)
	}
}

func TestEdgeCases(t *testing.T) {
	base10 := base_n.NewBaseN([]byte("0123456789"))

	testCases := []struct {
		name     string
		input    int64
		expected string
	}{
		{"Zero", 0, "0"},
		{"Negative", -42, "-42"},
		{"Max int64", 9223372036854775807, "9223372036854775807"},
		{"Min int64", -9223372036854775808, "-9223372036854775808"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := base10.Encode(tc.input)
			decoded, err := base10.Decode(encoded)

			if err != nil {
				t.Fatalf("Error decoding %s: %v", encoded, err)
			}

			if decoded != tc.input {
				t.Errorf("Expected decoded value to be %d, got %d", tc.input, decoded)
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	base10 := base_n.NewBaseN([]byte("0123456789"))

	_, err := base10.Decode("")
	if err == nil {
		t.Error("Expected error for empty string, got nil")
	}

	_, err = base10.Decode("1A")
	if err == nil {
		t.Error("Expected error for invalid character, got nil")
	}
}
