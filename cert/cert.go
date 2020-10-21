package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret"
	"io/ioutil"
	"math/big"
)

func ParseEnv(e string) string {
	return e[len("CERT_"):]
}

func GenerateCertificate() (ca []byte, key []byte, err error) {
	certificate := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
	}
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}

	caPrivateKeyPEM := new(bytes.Buffer)
	err = pem.Encode(caPrivateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey),
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	ca = caPEM.Bytes()
	key = caPrivateKeyPEM.Bytes()
	return
}

func WriteCertificate(certificateName string) (err error) {
	caFileName, keyFileName := CertificateFileName(certificateName)
	caPath := secret.KeyFilePath(caFileName)
	keyPath := secret.KeyFilePath(keyFileName)
	ca, key, err := GenerateCertificate()
	if err != nil {
		return
	}
	err = ioutil.WriteFile(caPath, ca, 0600)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = ioutil.WriteFile(keyPath, key, 0600)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CertificateFileName(certificateName string) (string, string) {
	return fmt.Sprintf("%s.cert", certificateName), fmt.Sprintf("%s.key", certificateName)
}

func LoadCertificate(certificateName string) (err error) {
	caFileName, _ := CertificateFileName(certificateName)
	if !secret.CheckSecretFileValid(caFileName) {
		err = WriteCertificate(certificateName)
	}
	return
}
