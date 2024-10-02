package ibkr

import (
	"crypto"
	"crypto/dsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ImportRsaKeyFromPem(pemPath string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ImportDhParametersFromPem(pemPath string) (*dsa.Parameters, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "DH PARAMETERS" {
		return nil, err
	}

	// Parse the DH parameters
	dhParams, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	dhParameters, ok := dhParams.(*dsa.Parameters)
	if !ok {
		return nil, fmt.Errorf("failed to cast dh params to *dsa.Parameters")
	}

	return dhParameters, nil
}

func SignRsa(input []byte, key *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()
	hash.Write(input)
	hashedMessage := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashedMessage)
}
