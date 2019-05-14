package paymentwall

const (
	VERSION = "0.0.1"

	APIBaseURL = "https://api.paymentwall.com/api"

	// API types
	APIVirtualCurrency = 1
	APIDigitalGoods    = 2
	APICart            = 3
)

var availableTypes = []int{
	APIVirtualCurrency,
	APIDigitalGoods,
	APICart,
}

func isValidAPIType(t int) bool {
	for _, v := range availableTypes {
		if t == v {
			return true
		}
	}

	return false
}

func signatureVersionSupported(v int) bool {
	if v < 1 || v > 3 {
		return false
	}

	return true
}
