package paymentwall

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

func CalculateSign(
	key string,
	params map[string]string,
	version int) string {
	hasher := getHasher(version)
	stringToHash := getStringToHash(key, params)

	hasher.Write([]byte(stringToHash))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getHasher(version int) hash.Hash {
	var hasher hash.Hash

	switch version {
	case 1, 2:
		hasher = md5.New()
		break

	case 3:
		hasher = sha256.New()
	}

	return hasher
}

func getStringToHash(key string, params map[string]string) string {
	params = SortParameters(params)
	output := []string{}

	for k, v := range params {
		output = append(output, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(output, "&")
}
