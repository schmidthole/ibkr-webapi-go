package ibkr

// import (
// 	"math/big"
// 	"reflect"
// 	"testing"
// )

// func TestBigToSignedBytes(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    *big.Int
// 		expected []byte
// 	}{
// 		{
// 			name:     "zero",
// 			input:    big.NewInt(0),
// 			expected: []byte{0x00},
// 		},
// 		{
// 			name:     "positive number 1",
// 			input:    big.NewInt(1),
// 			expected: []byte{0x00, 0x01},
// 		},
// 		// {
// 		// 	name:     "positive number 127 (boundary for single byte)",
// 		// 	input:    big.NewInt(127),
// 		// 	expected: []byte{0x7f},
// 		// },
// 		// {
// 		// 	name:     "positive number 128 (requires two bytes)",
// 		// 	input:    big.NewInt(128),
// 		// 	expected: []byte{0x00, 0x80},
// 		// },
// 		// {
// 		// 	name:     "negative number -1",
// 		// 	input:    big.NewInt(-1),
// 		// 	expected: []byte{0x80},
// 		// },
// 		// {
// 		// 	name:     "negative number -128 (boundary for negative single byte)",
// 		// 	input:    big.NewInt(-128),
// 		// 	expected: []byte{0x80},
// 		// },
// 		// {
// 		// 	name:     "negative number -129 (requires two bytes)",
// 		// 	input:    big.NewInt(-129),
// 		// 	expected: []byte{0x80, 0x81},
// 		// },
// 		// {
// 		// 	name:     "positive large number",
// 		// 	input:    new(big.Int).SetBytes([]byte{0x12, 0x34, 0x56}),
// 		// 	expected: []byte{0x12, 0x34, 0x56},
// 		// },
// 		// {
// 		// 	name:     "negative large number",
// 		// 	input:    new(big.Int).Neg(new(big.Int).SetBytes([]byte{0x12, 0x34, 0x56})),
// 		// 	expected: []byte{0x92, 0x34, 0x56}, // 0x12 flipped to 0x92 for sign
// 		// },
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := BigToSignedBytes(tt.input)
// 			if !reflect.DeepEqual(got, tt.expected) {
// 				t.Errorf("BigToSignedBytes(%v) = %v, want %v", tt.input, got, tt.expected)
// 			}
// 		})
// 	}
// }
