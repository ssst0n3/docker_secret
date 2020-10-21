package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret"
	"io/ioutil"
	"math/big"
	"time"
)

func ParseEnv(e string) string {
	return e[len("CERT_"):]
}

func CreateCertificateAuthority() (*x509.Certificate, *rsa.PrivateKey) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2020),
		IsCA: true,
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		awesome_error.CheckErr(err)
		return nil, nil
	}
	return ca, caPrivateKey
}

func GenerateCertificate() (cert []byte, key []byte, err error) {
	ca, caPrivateKey := CreateCertificateAuthority()
	certificate := &x509.Certificate{
		SerialNumber: big.NewInt(2020),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, ca, &privateKey.PublicKey, caPrivateKey)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	certificatePEM := new(bytes.Buffer)
	err = pem.Encode(certificatePEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificateBytes,
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}

	privateKeyPEM := new(bytes.Buffer)
	err = pem.Encode(privateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	cert = certificatePEM.Bytes()
	key = privateKeyPEM.Bytes()
	return
}

func WriteCertificate(certificateName string) (err error) {
	caPath, keyPath := CertificateFilePath(certificateName)
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
	return fmt.Sprintf("%s.crt", certificateName), fmt.Sprintf("%s.key", certificateName)
}

func CertificateFilePath(certificateName string) (string, string) {
	caFileName, keyFileName := CertificateFileName(certificateName)
	return secret.KeyFilePath(caFileName), secret.KeyFilePath(keyFileName)
}

func LoadCertificate(certificateName string) (err error) {
	caFileName, _ := CertificateFileName(certificateName)
	if !secret.CheckSecretFileValid(caFileName) {
		err = WriteCertificate(certificateName)
	}
	return
}
