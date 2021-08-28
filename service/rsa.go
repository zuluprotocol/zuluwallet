package service

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type RSAKeys struct {
	Pub  []byte
	Priv []byte
}

func GenerateRSAKeys() (*RSAKeys, error) {
	bitSize := 4096

	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, fmt.Errorf("unable to generate rsa keys: %w", err)
	}

	privateKey, err := toPrivatePKCS1Key(key)
	if err != nil {
		return nil, err
	}

	publicKey, err := toPublicPKCS1Key(key)
	if err != nil {
		return nil, err
	}

	return &RSAKeys{
		Pub:  publicKey,
		Priv: privateKey,
	}, nil
}

func toPrivatePKCS1Key(key *rsa.PrivateKey) ([]byte, error) {
	privateKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	privateKeyBuffer := bytes.NewBuffer([]byte{})
	err := pem.Encode(privateKeyBuffer, privateKey)
	if err != nil {
		return nil, err
	}
	return privateKeyBuffer.Bytes(), nil
}

func toPublicPKCS1Key(key *rsa.PrivateKey) ([]byte, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}
	publicKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	publicKeyBuffer := bytes.NewBuffer([]byte{})
	err = pem.Encode(publicKeyBuffer, publicKey)
	if err != nil {
		return nil, err
	}
	return publicKeyBuffer.Bytes(), nil
}
