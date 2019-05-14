package paymentwall_test

import (
	"reflect"
	"testing"

	paymentwall "github.com/paymentwall/paymentwall-go"
)

func TestSortParameters(t *testing.T) {
	tests := []struct {
		name string
		args map[string]string
		want map[string]string
	}{
		{
			name: "first case",
			args: map[string]string{
				"hello":   "world",
				"api":     "payment",
				"invoice": "api",
			},
			want: map[string]string{
				"api":     "payment",
				"hello":   "world",
				"invoice": "api",
			},
		},

		{
			name: "2nd test case",
			args: map[string]string{
				"a":     "quick",
				"brown": "fox",
				"jumps": "over",
				"the":   "lazy",
				"dog":   ".",
			},
			want: map[string]string{
				"a":     "quick",
				"brown": "fox",
				"dog":   ".",
				"jumps": "over",
				"the":   "lazy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paymentwall.SortParameters(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}
