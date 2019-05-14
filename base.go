package paymentwall

type app struct {
	cfg config
}

func New(config config) *app {
	return &app{
		cfg: config,
	}
}

func (a *app) NewWidget(code string, userID string, products []product) *widget {

	//validate
	if a.cfg.GetAPIType() == APIDigitalGoods && len(products) > 1 {
		panic("Digital Goods API does not support more than 1 product")
	}

	return &widget{
		cfg:              a.cfg,
		widgetAttributes: defaultWidgetAttributes(),

		Code:     code,
		UserID:   userID,
		Products: products,
	}
}
