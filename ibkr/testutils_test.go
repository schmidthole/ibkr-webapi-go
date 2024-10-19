package ibkr

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

type MockOAuthContext struct{}

func (i *MockOAuthContext) GenerateLiveSessionToken(client *http.Client, baseUrl string) error {
	return nil
}
func (i *MockOAuthContext) GetOAuthHeader(method string, requestUrl string) (string, error) {
	return "OAuth MOCK HEADER", nil
}
func (i *MockOAuthContext) ShouldReAuthenticate() bool {
	return false
}
func (i *MockOAuthContext) Reset() {}

func NewTestIbkrClient(baseUrl string) *IbkrWebClient {
	return &IbkrWebClient{
		BaseUrl: baseUrl,
		client:  &http.Client{Timeout: 15 * time.Second},
		oauth:   &MockOAuthContext{},
	}
}

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	exitCode := m.Run()

	log.SetOutput(os.Stderr)

	os.Exit(exitCode)
}
