package main

import (
	"reflect"
	"testing"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Example 1: Hello World",
			input: []string{"Hello", "World"},
		},
		{
			name:  "Example 2: empty string",
			input: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var codec Codec
			encoded := codec.Encode(tt.input)
			decoded := codec.Decode(encoded)
			if !reflect.DeepEqual(decoded, tt.input) {
				t.Errorf("Codec roundtrip failed: got %v, want %v", decoded, tt.input)
			}
		})
	}
}
