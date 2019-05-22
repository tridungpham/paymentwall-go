package paymentwall_test

import (
	"reflect"
	"testing"

	paymentwall "github.com/paymentwall/paymentwall-go"
)

func Test_SortParameters(t *testing.T) {
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

func Test_mergeMaps(t *testing.T) {
	type args struct {
		dest map[string]string
		src  map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			args: args{
				src: map[string]string{
					"hello": "world",
				},
				dest: map[string]string{
					"key": "value",
				},
			},
			want: map[string]string{
				"hello": "world",
				"key":   "value",
			},
		},

		{
			args: args{
				src: map[string]string{
					"hello": "world",
					"key":   "test",
				},
				dest: map[string]string{
					"key":  "value",
					"key2": "value2",
				},
			},
			want: map[string]string{
				"hello": "world",
				"key":   "test",
				"key2":  "value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paymentwall.MergeMaps(tt.args.dest, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapToQueryString(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				m: map[string]string{
					"abc": "10",
					"def": "20",
				},
			},
			want: "abc=10&def=20",
		},
		{
			args: args{
				m: map[string]string{
					"abc":       "10",
					"def":       "20",
					"map[key1]": "value1",
				},
			},
			want: "abc=10&def=20&map%5Bkey1%5D=value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paymentwall.MapToQueryString(tt.args.m); got != tt.want {
				t.Errorf("mapToQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
