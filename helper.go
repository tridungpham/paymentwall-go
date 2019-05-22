package paymentwall

import (
	"net/url"
	"sort"
)

func SortParameters(params map[string]string) map[string]string {
	var keys []string
	output := map[string]string{}

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		output[k] = params[k]
	}

	return output
}

func MergeMaps(dest, src map[string]string) map[string]string {
	for k := range src {
		dest[k] = src[k]
	}

	return dest
}

func MapToQueryString(m map[string]string) string {
	values := url.Values{}
	for k, v := range m {
		values.Add(k, v)
	}

	return values.Encode()
}
