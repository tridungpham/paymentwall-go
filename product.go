package paymentwall

const (
	SubscriptionProduct = "subscription"
	FixedProduct        = "fixed"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Currency string
	Type     string
}
