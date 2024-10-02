package ibkr

import (
	"fmt"
	"net/url"
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
		outString = outString + fmt.Sprintf("&%v=%v", key, p[key])
	}

	return url.PathEscape(outString)
}

func (p OAuthParams) ToHeaderString() string {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	outString := "OAuth "

	for _, key := range keys {
		outString = outString + fmt.Sprintf(", %v=%v", key, p[key])
	}

	return url.PathEscape(outString)
}
