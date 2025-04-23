package signutils

// These types just indicate the intended use case for
// a variable of the dedicated type.
// The values are interpreted by the appropriate GetXXX
// functions, which typically accept various
// kinds of instances:
//  - dedicated implementations
//  - byte sequences
//  - strings
//  - other elements or list of elements, which can be mapped appropriately.

// GenericPublicKey can be everything somebody
// can map to an appropriate PublicKey.
type GenericPublicKey interface{}

// GenericPrivateKey can be everything somebody
// can map to an appropriate PrivateKey.
type GenericPrivateKey interface{}

// GenericCertificate can be everything mappable
// by GetCertificate to an appropriate x509.Certificate.
type GenericCertificate interface{}

// GenericCertificateChain can be everything mappable
// by GetCertificateChain to an appropriate list of x509.Certificates.
// GenericCertificateChain is always a GenericCertificatePool.
type GenericCertificateChain interface{}

// GenericCertificatePool can be everything mappable
// by GetCertPool to an appropriate x509.CertPool.
type GenericCertificatePool interface{}

const (
	KIND_PUBLIC_KEY  = "public key"
	KIND_PRIVATE_KEY = "private key"
	KIND_CERTIFICATE = "certificate"
	KIND_CERTPOOL    = "certificate pool"
)
