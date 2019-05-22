package paymentwall_test

import (
	"testing"

	"github.com/paymentwall/paymentwall-go"
)

func Test_CalculateSign(t *testing.T) {
	t.Run("test signature calculation version 1", func(t *testing.T) {
		key := "some-secret-key"
		params := map[string]string{
			"name": "alex",
			"age":  "30",
		}
		sign := paymentwall.CalculateSign(key, params, 1)
		expected := "3b944829fd7589ffed0f66760243e82b"

		if sign != expected {
			t.Errorf(
				"Expected the signature to be %s, got %s instead",
				expected,
				sign,
			)
		}

		params["uid"] = "user-1"
		expected2 := "f3151489ee36cfd467b7f56bd89e2b1c"
		sign2 := paymentwall.CalculateSign(key, params, 1)

		if sign2 != expected2 {
			t.Errorf("Expected %s, got %s instead", expected2, sign2)
		}
	})

}
