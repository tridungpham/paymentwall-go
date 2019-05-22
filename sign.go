package paymentwall

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func CalculateSign(
	key string,
	params map[string]string,
	version int) string {
	hasher := getHasher(version)
	stringToHash := getStringToHash(version, params)
	stringToHash += key

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

func getStringToHash(version int, params map[string]string) string {
	switch version {
	case 1:
		if uid, ok := params["uid"]; ok {
			return uid
		}

		return ""

	case 2, 3:
		params = SortParameters(params)
		return MapToQueryString(params)
	}

	return ""
}
