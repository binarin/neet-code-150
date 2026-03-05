package main

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{
			name:    "Example 1",
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want:    []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name:    "Example 2",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findItinerary(tt.tickets)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
