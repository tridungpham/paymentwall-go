package paymentwall

type app struct {
	config Config
}

func New(config Config) *app {
	return &app{
		config: config,
	}
}

func (a *app) NewWidget(code string, userID string, products []Product) *widget {
	return &widget{
		config:           a.config,
		widgetAttributes: defaultWidgetAttributes(),

		Code:     code,
		UserID:   userID,
		Products: products,
	}
}
