package paymentwall

import (
	"fmt"
	"strconv"
)

const (
	SubscriptionProduct = "subscription"
	FixedProduct        = "fixed"
)

type ProductBaseData struct {
	ID       string
	Name     string
	Amount   float64
	Currency string
}

// Product struct that represent a product in customer system
type product struct {
	ProductBaseData

	// subscription or fixed
	productType string

	// Subscription settings
	isRecurring bool
}

func (p *product) GetAmountString() string {
	return fmt.Sprintf("%0.2f", p.Amount)
}

func (p *product) GetProductType() string {
	return p.productType
}

func (p *product) IsRecurring() bool {
	return p.isRecurring
}

func (p product) GetDataMap() map[string]string {
	return map[string]string{
		"amount":         p.GetAmountString(),
		"currencyCode":   p.Currency,
		"ag_name":        p.Name,
		"ag_external_id": p.ID,
		"ag_type":        p.GetProductType(),
	}
}

type subscriptionProduct struct {
	*product
	*period

	trialProduct *subscriptionProduct
}

func (sp *subscriptionProduct) GetPeriodType() string {
	return sp.period.Type
}

func (sp *subscriptionProduct) GetPeriodLength() int {
	return sp.period.Length
}

func (sp *subscriptionProduct) SetTrialProduct(trial *subscriptionProduct) {
	if trial.HasTrialProduct() {
		panic("Trial product can not have another trial product")
	}

	sp.trialProduct = trial
}

func (sp *subscriptionProduct) HasTrialProduct() bool {
	return sp.trialProduct != nil
}

func (sp *subscriptionProduct) GetTrialProduct() *subscriptionProduct {
	return sp.trialProduct
}

func (sp *subscriptionProduct) GetDataMap() map[string]string {
	var base map[string]string

	if sp.HasTrialProduct() {
		base = sp.trialProduct.GetDataMap()

		base["ag_trial"] = "1"
		base["ag_post_trial_external_id"] = sp.ID
		base["ag_post_trial_period_length"] = strconv.Itoa(sp.GetPeriodLength())
		base["ag_post_trial_period_type"] = sp.GetPeriodType()
		base["ag_post_trial_name"] = sp.Name
		base["post_trial_amount"] = sp.GetAmountString()
		base["post_trial_currencyCode"] = sp.Currency
	} else {
		base = sp.product.GetDataMap()
		base["ag_period_type"] = sp.period.Type
		base["ag_period_length"] = strconv.Itoa(sp.GetPeriodLength())
		base["ag_recurring"] = "1"
	}

	return base
}

func NewFixedProduct(data ProductBaseData) *product {
	validateProductBaseData(data)

	return &product{
		ProductBaseData: data,

		productType: FixedProduct,
	}
}

func NewSubscriptionProduct(
	data ProductBaseData,
	periodConfig *period,
) *subscriptionProduct {
	validateProductBaseData(data)

	product := &subscriptionProduct{
		product: &product{
			ProductBaseData: data,

			productType: SubscriptionProduct,
			isRecurring: true,
		},
		period: periodConfig,
	}

	return product
}

func validateProductBaseData(data ProductBaseData) {
	if data.ID == "" {
		panic(fmt.Errorf("Product ID was not provided"))
	}

	if data.Name == "" {
		panic(fmt.Errorf("Product name was not provided"))
	}

	if data.Amount < 0 {
		panic(fmt.Errorf("Invalid price amount provided: %f", data.Amount))
	}

	if data.Currency == "" {
		panic(fmt.Errorf("Price currency was not provided"))
	}
}
