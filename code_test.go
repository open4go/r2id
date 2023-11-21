package r2id

import "testing"

func TestB6Code(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"test",
			198922,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := B6Code(); got != tt.want {
				t.Errorf("B6Code() = %v, want %v", got, tt.want)
			}
		})
	}
}
