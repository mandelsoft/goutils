package signutils

import (
	"crypto/x509"
	"crypto/x509/pkix"

	"github.com/mandelsoft/goutils/errors"
)

func VerifyCert(intermediate GenericCertificateChain, root GenericCertificatePool, name string, cert *x509.Certificate) error {
	return VerifyCertDN(intermediate, root, CommonName(name), cert)
}

func VerifyCertDN(intermediate GenericCertificateChain, root GenericCertificatePool, name *pkix.Name, cert *x509.Certificate) error {
	rootPool, err := GetCertPool(root, false)
	if err != nil {
		return err
	}
	interPool, err := GetCertPool(intermediate, false)
	if err != nil {
		return err
	}
	opts := x509.VerifyOptions{
		Intermediates: interPool,
		Roots:         rootPool,
		CurrentTime:   cert.NotBefore,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageCodeSigning},
	}
	_, err = cert.Verify(opts)
	if err != nil {
		return err
	}
	if name != nil {
		if err := MatchDN(cert.Subject, *name); err != nil {
			return err
		}
	}
	if cert.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
		return nil
	}
	for _, k := range cert.ExtKeyUsage {
		if k == x509.ExtKeyUsageCodeSigning {
			return nil
		}
	}
	return errors.ErrNotSupported("codesign", "", "certificate")
}
