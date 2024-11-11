package main

import (
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) []int
		inspect func([]int, *testing.T)
	}{
		{
			name: "sort a slice of integers",
			init: func(t *testing.T) []int {
				return []int{12, 3, 7, 9, 14, 6, 11, 2, 10, 5, 4, 8, 13, 1}
			},
			inspect: func(nums []int, t *testing.T) {
				merge_sort(nums)
				expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				if !reflect.DeepEqual(expected, nums) {
					t.Errorf("For input %v, expected %v", nums, expected)
				}
			},
		},	
		{
			name: "size of the slice",
			init: func(t *testing.T) []int {
				return []int{9, 6, 5, 3, 8, 0, 4, 2, 7, 1}
			},
			inspect: func(nums []int, t *testing.T) {			
				merge_sort(nums)
				expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				if len(nums) != len(expected) {
					t.Fatalf("Expected %v but got %v", len(nums), len(expected))
				}
			},

		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}
		})
	}
}