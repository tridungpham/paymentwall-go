package paymentwall

import (
	"fmt"
	"strconv"
)

type WidgetAttributes map[string]string

type widget struct {
	cfg config

	Code     string
	UserID   string
	Products []product

	//
	widgetAttributes WidgetAttributes
}

func (w *widget) SetConfig(config config) *widget {
	w.cfg = config
	return w
}

func (w *widget) SetHeight(height int) *widget {
	return w.SetAttribute("height", strconv.Itoa(height))
}

func (w *widget) SetWidth(width int) *widget {
	return w.SetAttribute("width", strconv.Itoa(width))
}

func (w *widget) SetAttribute(key, value string) *widget {
	w.widgetAttributes[key] = value
	return w
}

func (w *widget) SetProducts(products []product) {
	w.Products = products
}

func (w *widget) GetURL() string {
	return fmt.Sprintf(
		"%s/%s?%s",
		APIBaseURL,
		w.GetController(),
		w.BuildQuery(),
	)
}

func (w *widget) GetHTML() string {
	attributes := ""
	for key, value := range w.widgetAttributes {
		attributes += fmt.Sprintf("%s=\"%s\" ", key, value)
	}

	html := "<iframe src=\"%s\" %s ></iframe>"

	return fmt.Sprintf(html, w.GetURL(), attributes)
}

func (w *widget) GetController() string {
	switch w.cfg.GetAPIType() {
	case APIVirtualCurrency:
		return "ps"

	case APIDigitalGoods:
		return "subscription"

	case APICart:
		return "cart"

	default:
		return ""
	}
}

func (w *widget) BuildQuery() string {
	params := map[string]string{
		"key":          w.cfg.GetPublicKey(),
		"uid":          w.UserID,
		"widget":       w.Code,
		"sign_version": strconv.Itoa(w.cfg.GetSignatureVersion()),
	}

	switch w.cfg.GetAPIType() {
	case APIDigitalGoods:
		params = mergeMaps(
			params,
			prepareDigitalGoodsParameters(w.Products),
		)
		break

	case APICart:
		params = mergeMaps(
			params,
			prepareCartParameters(w.Products),
		)
		break
	}

	params["sign"] = CalculateSign(
		w.cfg.GetPrivateKey(),
		params,
		w.cfg.GetSignatureVersion(),
	)

	return mapToQueryString(params)
}

func defaultWidgetAttributes() WidgetAttributes {
	return WidgetAttributes{
		"width":       "800",
		"height":      "600",
		"frameborder": "0",
	}
}

func prepareDigitalGoodsParameters(products []product) map[string]string {
	if len(products) > 0 {
		// take only the first product
		product := products[0]

		return product.GetDataMap()
	}

	return map[string]string{}
}

func prepareCartParameters(products []product) map[string]string {
	return map[string]string{}
}
