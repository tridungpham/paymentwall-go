package paymentwall_test

import (
	"testing"

	"github.com/paymentwall/paymentwall-go"
)

func TestCalculateSign(t *testing.T) {
	type args struct {
		key     string
		params  map[string]string
		version int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test signature version 1",
			args: args{
				key: "some-secret-key",
				params: map[string]string{
					"name": "alex",
					"age":  "30",
				},
				version: 1,
			},
			want: "3b944829fd7589ffed0f66760243e82b",
		},
		{
			name: "test signature version 1 with user id",
			args: args{
				key: "some-secret-key",
				params: map[string]string{
					"name": "alex",
					"age":  "30",
					"uid":  "user-1",
				},
				version: 1,
			},
			want: "f3151489ee36cfd467b7f56bd89e2b1c",
		},
		{
			name: "test signature version 2",
			args: args{
				key: "some-secret-key",
				params: map[string]string{
					"name": "alex",
					"age":  "30",
				},
				version: 2,
			},
			want: "c4d3d531e1f87cd5155c02bd9703ec3c",
		},
		{
			name: "test signature version 3",
			args: args{
				key: "some-secret-key",
				params: map[string]string{
					"name": "alex",
					"age":  "30",
				},
				version: 3,
			},
			want: "fee0b06d951f2511bc6e6b86c0a30d44b178e074359ae7f3a366d87f6179bae4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paymentwall.CalculateSign(tt.args.key, tt.args.params, tt.args.version); got != tt.want {
				t.Errorf("CalculateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
