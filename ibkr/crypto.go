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
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, err
	}

	pkcs1Key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		return pkcs8Key.(*rsa.PrivateKey), err
	}

	return pkcs1Key, err
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
	hash := sha256.New()
	hash.Write(input)
	hashedMessage := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashedMessage)
}
