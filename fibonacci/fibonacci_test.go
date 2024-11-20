package main

import (
	"reflect"
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) int
		inspect func(int, *testing.T)
	}{
		{
			name: "Finbonnaci of 5",
			init: func(t *testing.T) int {
				return 6
			},
			inspect: func(num int, t *testing.T) {
				ret := fib(num)
				expected := 21
				if !reflect.DeepEqual(expected, ret) {
					t.Errorf("For input %v, expected %v", ret, expected)
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