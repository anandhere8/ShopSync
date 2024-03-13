package key

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

const (
	privateKeyPath = "key/private.key"
	publicKeyPath  = "key/public.key"
)

func loadECDSAPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	// Read the private key file
	keyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM encoded data
	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing EC private key")
	}

	// Parse the ECDSA private key
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func loadECDSAPublicKey(filename string) (*ecdsa.PublicKey, error) {
	// Read the public key file
	keyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM encoded data
	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the ECDSA public key
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Assert the parsed public key to be of type *ecdsa.PublicKey
	pubKey, ok := pubInterface.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not of type *ecdsa.PublicKey")
	}

	return pubKey, nil
}

func LoadPrivetKey() (*ecdsa.PrivateKey, error) {

	return loadECDSAPrivateKey(privateKeyPath)

}

func LoadPublicKey() (*ecdsa.PublicKey, error) {

	return loadECDSAPublicKey(publicKeyPath)

}
