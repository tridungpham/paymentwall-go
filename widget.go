package paymentwall

import (
	"fmt"
	"strconv"
)

type WidgetAttributes map[string]string

type widget struct {
	config Config

	Code     string
	UserID   string
	Products []Product

	//
	widgetAttributes WidgetAttributes
}

func (w *widget) SetConfig(config Config) *widget {
	w.config = config
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
	switch w.config.APIType {
	case APIVirtualCurrency:
		return "ps"

	case APIDigitalGoods:
		return "subscription"

	case APICart:
		return "cart"

	default:
		panic(fmt.Sprintf("Invalid API Type: %v", w.config.APIType))
	}
}

func (w *widget) BuildQuery() string {
	return ""
}

func defaultWidgetAttributes() WidgetAttributes {
	return WidgetAttributes{
		"width":       "800",
		"height":      "600",
		"frameborder": "0",
	}
}
