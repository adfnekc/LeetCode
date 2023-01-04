package main

import "testing"

func TestMinStack(t *testing.T) {
	type fields struct {
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"min_statck", fields{}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor_MinStack()
			obj.Push(-2)
			obj.Push(0)
			obj.Push(-3)
			obj.Pop()
			if got := obj.min(); got != tt.want {
				t.Errorf("MinStack.min() = %v, want %v", got, tt.want)
			}
		})
	}
}
