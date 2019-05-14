package paymentwall_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	paymentwall "github.com/paymentwall/paymentwall-go"
)

func Test_NewFixedProduct(t *testing.T) {
	productID := "product-id"
	productName := "acme-product"
	productSamplePrice := 9.99
	productCurrency := "USD"

	product := paymentwall.NewFixedProduct(
		paymentwall.ProductBaseData{
			ID:       productID,
			Name:     productName,
			Amount:   productSamplePrice,
			Currency: productCurrency,
		},
	)

	if product.Name != productName ||
		product.ID != productID ||
		product.Amount != productSamplePrice ||
		product.Currency != productCurrency {
		t.Error("Base attributes do not equal")
	}

	if product.GetProductType() != paymentwall.FixedProduct {
		t.Errorf(
			"Expect product's type is fixed, got '%s' instead",
			product.GetProductType(),
		)
	}

	if product.IsRecurring() == true {
		t.Error("Expect product is not recurring type")
	}

	if product.GetAmountString() != "9.99" {
		t.Errorf("Expect product amount is 9.99, got '%s' instead",
			product.GetAmountString(),
		)
	}
}

func Test_FixedProduct_GetDataMap(t *testing.T) {
	tests := []struct {
		name   string
		fields paymentwall.ProductBaseData
		want   map[string]string
	}{
		{
			name: "Basic case",
			fields: paymentwall.ProductBaseData{
				ID:       "sample-id",
				Name:     "sample-name",
				Amount:   9.99,
				Currency: "USD",
			},
			want: map[string]string{
				"amount":         "9.99",
				"currencyCode":   "USD",
				"ag_name":        "sample-name",
				"ag_external_id": "sample-id",
				"ag_type":        "fixed",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := paymentwall.NewFixedProduct(tt.fields)
			if got := p.GetDataMap(); !reflect.DeepEqual(
				paymentwall.SortParameters(got),
				paymentwall.SortParameters(tt.want),
			) {
				t.Errorf("product.GetDataMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewSubscriptionProduct(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		a := assert.New(t)

		productID := "product-id"
		productName := "acme-product"
		productSamplePrice := 9.99
		productCurrency := "USD"

		product := paymentwall.NewSubscriptionProduct(
			paymentwall.ProductBaseData{
				ID:       productID,
				Name:     productName,
				Amount:   productSamplePrice,
				Currency: productCurrency,
			},

			paymentwall.NewPeriod(
				paymentwall.PeriodTypeDay,
				30,
			),
		)

		a.Equal(productID, product.ID)
		a.Equal(productName, product.Name)
		a.Equal(productSamplePrice, product.Amount)
		a.Equal(productCurrency, product.Currency)
		a.Equal(paymentwall.SubscriptionProduct, product.GetProductType())
		a.True(product.IsRecurring())
		a.Equal(paymentwall.PeriodTypeDay, product.GetPeriodType())
		a.Equal(30, product.GetPeriodLength())
		a.False(product.HasTrialProduct())
		a.Nil(product.GetTrialProduct())
	})
}

func Test_SubscriptionProduct_GetDataMap(t *testing.T) {
	t.Run("subscription product without trial product", func(t *testing.T) {
		a := assert.New(t)

		productID := "product-id"
		productName := "acme-product"
		productSamplePrice := 9.99
		productCurrency := "USD"

		product := paymentwall.NewSubscriptionProduct(
			paymentwall.ProductBaseData{
				ID:       productID,
				Name:     productName,
				Amount:   productSamplePrice,
				Currency: productCurrency,
			},

			paymentwall.NewPeriod(
				paymentwall.PeriodTypeDay,
				30,
			),
		)

		expected := map[string]string{
			"amount":           "9.99",
			"currencyCode":     "USD",
			"ag_name":          "acme-product",
			"ag_external_id":   "product-id",
			"ag_type":          "subscription",
			"ag_period_length": "30",
			"ag_period_type":   "day",
			"ag_recurring":     "1",
		}

		a.Equal(
			paymentwall.SortParameters(expected),
			paymentwall.SortParameters(product.GetDataMap()),
		)
	})

	t.Run("subscription product with trial", func(t *testing.T) {
		a := assert.New(t)

		productID := "product-id"
		productName := "acme-product"
		productSamplePrice := 9.99
		productCurrency := "USD"

		product := paymentwall.NewSubscriptionProduct(
			paymentwall.ProductBaseData{
				ID:       productID,
				Name:     productName,
				Amount:   productSamplePrice,
				Currency: productCurrency,
			},

			paymentwall.NewPeriod(
				paymentwall.PeriodTypeWeek,
				2,
			),
		)

		trialProductID := "trial-product-id"
		trialProductName := "acme-trial-product"
		trialProductPrice := 1.0
		trialProductCurrency := "VND"

		trialProduct := paymentwall.NewSubscriptionProduct(
			paymentwall.ProductBaseData{
				ID:       trialProductID,
				Name:     trialProductName,
				Amount:   trialProductPrice,
				Currency: trialProductCurrency,
			},
			paymentwall.NewPeriod(
				paymentwall.PeriodTypeDay,
				3,
			),
		)

		product.SetTrialProduct(trialProduct)

		expected := map[string]string{
			"amount":           "1.00",
			"currencyCode":     trialProductCurrency,
			"ag_name":          trialProductName,
			"ag_external_id":   trialProductID,
			"ag_type":          paymentwall.SubscriptionProduct,
			"ag_period_length": "3",
			"ag_period_type":   paymentwall.PeriodTypeDay,
			"ag_recurring":     "1",
			"ag_trial":         "1",

			"ag_post_trial_external_id":   productID,
			"ag_post_trial_name":          productName,
			"ag_post_trial_period_length": "2",
			"ag_post_trial_period_type":   paymentwall.PeriodTypeWeek,
			"post_trial_amount":           "9.99",
			"post_trial_currencyCode":     "USD",
		}

		a.Equal(
			paymentwall.SortParameters(expected),
			paymentwall.SortParameters(product.GetDataMap()),
		)
	})
}
