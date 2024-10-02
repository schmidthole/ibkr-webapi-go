package ibkr

import (
	"net/url"
	"testing"
)

func unescape(s string) string {
	res, _ := url.PathUnescape(s)
	return res
}

func TestOAuthParams_ToSignatureString(t *testing.T) {
	tests := []struct {
		name     string
		input    OAuthParams
		expected string
	}{
		{
			name: "normal case",
			input: OAuthParams{
				"oauth_consumer_key":     "key123",
				"oauth_token":            "token123",
				"oauth_signature_method": "HMAC-SHA1",
				"oauth_timestamp":        "123456789",
			},
			expected: "&oauth_consumer_key=key123&oauth_signature_method=HMAC-SHA1&oauth_timestamp=123456789&oauth_token=token123",
		},
		{
			name:     "empty params",
			input:    OAuthParams{},
			expected: "",
		},
		{
			name: "special characters",
			input: OAuthParams{
				"oauth_token":        "token$@#",
				"oauth_consumer_key": "key!?",
			},
			expected: "&oauth_consumer_key=key!?&oauth_token=token$@#",
		},
		{
			name: "single parameter",
			input: OAuthParams{
				"oauth_token": "token123",
			},
			expected: "&oauth_token=token123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := unescape(tt.input.ToSignatureString())
			if result != tt.expected {
				t.Errorf("ToSignatureString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestOAuthParams_ToHeaderString(t *testing.T) {
	tests := []struct {
		name     string
		input    OAuthParams
		expected string
	}{
		{
			name: "normal case",
			input: OAuthParams{
				"oauth_consumer_key":     "key123",
				"oauth_token":            "token123",
				"oauth_signature_method": "HMAC-SHA1",
				"oauth_timestamp":        "123456789",
			},
			expected: "OAuth , oauth_consumer_key=key123, oauth_signature_method=HMAC-SHA1, oauth_timestamp=123456789, oauth_token=token123",
		},
		{
			name:     "empty params",
			input:    OAuthParams{},
			expected: "OAuth ",
		},
		{
			name: "special characters",
			input: OAuthParams{
				"oauth_token":        "token$@#",
				"oauth_consumer_key": "key!?",
			},
			expected: "OAuth , oauth_consumer_key=key!?, oauth_token=token$@#",
		},
		{
			name: "single parameter",
			input: OAuthParams{
				"oauth_token": "token123",
			},
			expected: "OAuth , oauth_token=token123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := unescape(tt.input.ToHeaderString())
			if result != tt.expected {
				t.Errorf("ToHeaderString() = %v, want %v", result, tt.expected)
			}
		})
	}
}
