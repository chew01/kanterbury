package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

// Helper function to create a x509 certificate template
func certTemplate() (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("failed to generate serial number: " + err.Error())
	}

	issuerName := pkix.Name{
		Organization:       []string{"Kanterbury"},
		OrganizationalUnit: []string{"Kanterbury"},
		CommonName:         "github.com/chew01/kanterbury",
	}
	template := x509.Certificate{
		BasicConstraintsValid: true,
		Issuer:                issuerName,
		NotAfter:              time.Now().AddDate(5, 0, 0), // 5 years expiration for CA
		NotBefore:             time.Now(),
		SerialNumber:          serialNumber,
		SignatureAlgorithm:    x509.SHA256WithRSA,
		Subject:               issuerName,
	}

	return &template, nil
}

// Helper function to create certificate, and in PEM format
func createCert(template, parent *x509.Certificate, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) (cert *x509.Certificate, certPEM []byte, err error) {
	certDER, err := x509.CreateCertificate(rand.Reader, template, parent, publicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	cert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, errors.New("failed to parse certificate: " + err.Error())
	}

	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM = pem.EncodeToMemory(&b)
	return cert, certPEM, nil
}

// GenerateCA generates a new key pair and saves at the given paths
func GenerateCA(certPath, keyPath string) error {
	rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("error generating rsa key: %v", err)
	}

	rootCertTmpl, err := certTemplate()
	if err != nil {
		return fmt.Errorf("error generating cert template: %v", err)
	}

	rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature
	rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
	rootCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}

	_, rootCertPEM, err := createCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
	if err != nil {
		return fmt.Errorf("error creating cert: %v", err)
	}
	rootKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey)})

	Must(os.WriteFile(certPath, rootCertPEM, 0666))
	Must(os.WriteFile(keyPath, rootKeyPEM, 0666))

	return nil
}
