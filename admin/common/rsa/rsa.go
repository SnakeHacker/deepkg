package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/golang/glog"
)

// GenerateKey Generates the RSA key
func GenerateKey(bits int) (privateKey *rsa.PrivateKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

// Encrypt enctypts the string data
func Encrypt(data string, publicKey *rsa.PublicKey) (encryptedData string, err error) {
	dataBytes := []byte(data)

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, dataBytes)
	if err != nil {
		glog.Error(err)
		return
	}

	encryptedData = base64.StdEncoding.EncodeToString(encryptedBytes)

	return
}

// Decrypt decrypts the string data
func Decrypt(encryptedData string, privateKey *rsa.PrivateKey) (decryptedData string, err error) {
	dataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		glog.Error(err)
		return
	}

	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, dataBytes)
	if err != nil {
		glog.Error(err)
		return
	}

	decryptedData = string(decryptedBytes)

	return
}

func LoadPublicKey(pemData string) (publicKey *rsa.PublicKey, err error) {
	pubPEMData := []byte(fmt.Sprintf(`
-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----
	`, pemData))

	block, _ := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		err = errors.New("failed to decode PEM block containing public key")
		glog.Error(err)
		return
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		glog.Error(err)
		return
	}

	return pubKey.(*rsa.PublicKey), nil

}
