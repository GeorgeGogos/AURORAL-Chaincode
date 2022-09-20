package test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/s7techlab/cckit/identity"
)

var (
	caPrivateKey *ecdsa.PrivateKey
)

func GenerateSelfSignedPEMCertBytes(commonName, orgID string) ([]byte, error) {
	if caPrivateKey == nil {
		priv, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
		if err != nil {
			return nil, fmt.Errorf("Error while generating CA private key: %s", err.Error())
		}
		caPrivateKey = priv
	}
	keyUsage := x509.KeyUsageDigitalSignature
	notBefore := time.Now()
	validFor := 150000 * time.Second
	notAfter := notBefore.Add(validFor)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	certificateTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
			CommonName:   commonName,
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certificateTemplate.IPAddresses = append(certificateTemplate.IPAddresses, net.ParseIP("127.0.0.1"))
	certificateTemplate.IsCA = false

	//Inserting custom X509 Extension into cert
	buf, _ := json.Marshal(map[string]string{"cid": orgID})
	ext := pkix.Extension{
		Id:       asn1.ObjectIdentifier{1, 2, 3, 4, 5, 6, 7, 8, 1},
		Critical: false,
		Value:    buf,
	}
	certificateTemplate.ExtraExtensions = []pkix.Extension{ext}

	certificateDERBytes, err := x509.CreateCertificate(rand.Reader, &certificateTemplate,
		&certificateTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("Error while creating DER-encoded X.509 digital certificate: %s", err.Error())
	}

	certBuffer := bytes.NewBuffer(nil)
	err = pem.Encode(certBuffer, &pem.Block{Type: "CERTIFICATE", Bytes: certificateDERBytes})
	if err != nil {
		return nil, fmt.Errorf("Error while PEM encoding X.509 digital certificate: %s", err.Error())
	}
	return certBuffer.Bytes(), nil
}

func GenerateCertIdentity(mspID, commonName, orgID string) (*identity.CertIdentity, error) {
	certPEMBytes, err := GenerateSelfSignedPEMCertBytes(commonName, orgID)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(certPEMBytes))
	return identity.New(mspID, certPEMBytes)
}
