package test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
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

type AttributeTest struct {
	Name, Value string
}

func (a *AttributeTest) GetName() string {
	return a.Name
}

func (a *AttributeTest) GetValue() string {
	return a.Value
}

type AttributeRequestTest struct {
	Name    string
	Require bool
}

func (ar *AttributeRequestTest) GetName() string {
	return ar.Name
}

func (ar *AttributeRequestTest) IsRequired() bool {
	return ar.Require
}

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
	mgr := New()
	attrs := []Attribute{
		&AttributeTest{Name: "cid", Value: orgID},
	}
	reqs := []AttributeRequest{
		&AttributeRequestTest{Name: "cid", Require: true},
	}
	err = mgr.ProcessAttributeRequestsForCert(reqs, attrs, &certificateTemplate)
	if err != nil {
		retErr := fmt.Errorf("Failed to ProcessAttributeRequestsForCert: %s", err)
		return nil, retErr
	}
	//marshaledCert, _ := json.Marshal(certificateTemplate)
	//fmt.Printf("CERT RAW IS: %s", string(marshaledCert))

	//End of custom X509 Extension insertion

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
	fmt.Print(string(certBuffer.Bytes()))
	return certBuffer.Bytes(), nil
}

func GenerateCertIdentity(mspID, commonName, orgID string) (*identity.CertIdentity, error) {
	certPEMBytes, err := GenerateSelfSignedPEMCertBytes(commonName, orgID)
	if err != nil {
		return nil, err
	}

	return identity.New(mspID, certPEMBytes)
}
