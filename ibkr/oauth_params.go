package ibkr

import (
	"fmt"
	"log"
	"net/url"
	"sort"
)

type OAuthParams map[string]string

func (p OAuthParams) logRaw() {
	log.Println("oauth raw params: ")
	for key, val := range p {
		log.Printf("- %v: %v", key, val)
	}
}

func (p OAuthParams) ToSignatureString() string {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	outString := ""

	for i, key := range keys {
		if i == 0 {
			outString = outString + fmt.Sprintf("%v=%v", key, p[key])
		} else {
			outString = outString + fmt.Sprintf("&%v=%v", key, p[key])
		}

	}

	return "&" + url.QueryEscape(outString)
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
