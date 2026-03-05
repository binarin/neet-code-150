package main

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
	}
	for _, tt := range tests {
		got := partitionLabels(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("partitionLabels(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
