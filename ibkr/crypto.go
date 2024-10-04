package ibkr

import (
	"crypto"
	"crypto/dsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

type dhParams struct {
	P *big.Int
	G *big.Int
}

func ImportRsaKeyFromPem(pemPath string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)

	pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := pkcs8Key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("error parsing PKCS#8 rsa key")
	}

	return rsaKey, nil
}

func ImportDhParametersFromPem(pemPath string) (*dsa.Parameters, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "DH PARAMETERS" {
		return nil, fmt.Errorf("invalid PEM block type: %s", block.Type)
	}

	var params dhParams
	_, err = asn1.Unmarshal(block.Bytes, &params)
	if err != nil {
		return nil, err
	}

	return &dsa.Parameters{
		P: params.P,
		Q: big.NewInt(0),
		G: params.G,
	}, nil
}

func SignRsa(input []byte, key *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(input)
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash[:])
}
