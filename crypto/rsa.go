package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func RsaKeyGen(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	priFile, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer priFile.Close()
	err = pem.Encode(priFile, &block)
	if err != nil {
		return err
	}

	publicKey := privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil
	}
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	pubFile, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer pubFile.Close()
	err = pem.Encode(pubFile, &block)
	if err != nil {
		return err
	}
	return nil

}
