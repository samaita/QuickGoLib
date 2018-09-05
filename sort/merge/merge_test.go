package merge

import (
	"reflect"
	"testing"
)

func TestSortV1(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "Sort ASC",
			input: []int64{8, 9, 2, 3, 1},
			want:  []int64{9, 8, 3, 2, 1},
		},
		{
			name:  "Sort ASC 2",
			input: []int64{8, 9, 2, 5, 3, 1},
			want:  []int64{9, 8, 5, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortV1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortV2(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "Sort ASC",
			input: []int64{8, 9, 2, 3, 1},
			want:  []int64{9, 8, 3, 2, 1},
		},
		{
			name:  "Sort ASC 2",
			input: []int64{8, 9, 2, 5, 3, 1},
			want:  []int64{9, 8, 5, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortV2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
