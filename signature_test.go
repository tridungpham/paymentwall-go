package paymentwall

import "testing"

func TestCalculateSignature(t *testing.T) {
	privateKey := "my-private-key"
	data := make(map[string]interface{})

	data["key1"] = "Hello"
	data["key2"] = 2
	data["key3"] = map[string]string{"key4": "hello", "key5": "world", "abc": "def"}
	data["key0"] = "yolo"
	data["key11"] = nil
	data["key111"] = 1.2

	correct := "131a744cdde8eaf40ee3bfa9c821f748fb29a68c00b18adc41964e3748bf010b"
	sign := CalculateSignature(privateKey, data, 3)
	if sign != correct {
		t.Error("Expected: ", correct, " got ", sign)
	}
}