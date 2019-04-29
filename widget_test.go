package paymentwall

import (
	"crypto/md5"
	"testing"
)

func Test_widget_GetController(t *testing.T) {
	tests := []struct {
		name    string
		apiType int
		want    string
	}{
		{
			apiType: APIVirtualCurrency,
			want:    "ps",
		},
		{
			apiType: APIDigitalGoods,
			want:    "subscription",
		},
		{
			apiType: APICart,
			want:    "cart",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &widget{
				config: Config{
					APIType: tt.apiType,
				},
			}
			if got := w.GetController(); got != tt.want {
				t.Errorf("When API Type is '%v': widget.GetController() = '%v', want '%v'", tt.apiType, got, tt.want)
			}
		})
	}

	t.Run("should panic when providing invalid API type", func(t *testing.T) {
		defer func() {
			r := recover()
			expectedPanicInfo := "Invalid API Type: 8"
			if r != expectedPanicInfo {
				t.Errorf("Expect panic with error description: '%s'", expectedPanicInfo)
			}
		}()

		w := &widget{config: Config{APIType: 8}}
		w.GetController()
	})
}

func Test_widget_BuildQuery(t *testing.T) {
	publicKey := md5.Sum([]byte("public"))
	privateKey := md5.Sum([]byte("private"))

	config := Config{
		APIType:    APIVirtualCurrency,
		PublicKey:  string(publicKey[:]),
		PrivateKey: string(privateKey[:]),
	}

	app := New(config)

	tests := []struct {
		name string
		w    *widget
		want string
	}{
		{
			w: app.NewWidget("pw", "user1", []Product{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.BuildQuery(); got != tt.want {
				t.Errorf("widget.BuildQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
