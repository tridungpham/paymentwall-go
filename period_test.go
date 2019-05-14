package paymentwall_test

import (
	"testing"

	pw "github.com/paymentwall/paymentwall-go"
	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	t.Run("should panic for invalid period type", func(t *testing.T) {
		defer func() {
			r := recover()
			err, ok := r.(error)
			if r == nil || !ok {
				t.Error("Should panic when provide invalid period type")
			}

			if err.Error() != "Invalid period type: 'invalid-period-type'" {
				t.Errorf("Invalid panic message: %s", r)
			}
		}()

		pw.NewPeriod("invalid-period-type", 30)
	})

	t.Run("success case", func(t *testing.T) {
		p := pw.NewPeriod(
			pw.PeriodTypeDay,
			30,
		)

		assert.Equal(t, pw.PeriodTypeDay, p.Type)
		assert.Equal(t, 30, p.Length)
	})
}
