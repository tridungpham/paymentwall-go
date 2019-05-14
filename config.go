package paymentwall

import "fmt"

type config struct {
	apiType    int
	publicKey  string
	privateKey string

	signVersion int
}

func NewConfig(
	apiType int,
	publicKey string,
	privateKey string,
	signVersion int,
) *config {
	if isValidAPIType(apiType) == false {
		panic(fmt.Errorf("Invalid API Type provided: %s", apiType))
	}

	if !signatureVersionSupported(signVersion) {
		panic(fmt.Errorf("Signature version %d is not supported", signVersion))
	}

	cfg := &config{
		apiType:     apiType,
		publicKey:   publicKey,
		privateKey:  privateKey,
		signVersion: signVersion,
	}

	return cfg
}

func (c config) GetAPIType() int {
	return c.apiType
}

func (c config) GetPublicKey() string {
	return c.publicKey
}

func (c config) GetPrivateKey() string {
	return c.privateKey
}

func (c config) GetSignatureVersion() int {
	return c.signVersion
}
