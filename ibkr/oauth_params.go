package ibkr

import (
	"fmt"
	"sort"
)

type OAuthParams map[string]string

func (p OAuthParams) ToSignatureString() string {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	outString := ""

	for _, key := range keys {
		outString = outString + fmt.Sprintf("&%v=\"%v\"", key, p[key])
	}

	return outString
}

func (p OAuthParams) ToHeaderString() string {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	outString := "OAuth "

	for i, key := range keys {
		if i == 0 {
			outString = outString + fmt.Sprintf("%v=\"%v\"", key, p[key])
		} else {
			outString = outString + fmt.Sprintf(", %v=\"%v\"", key, p[key])
		}
	}

	return outString
}
