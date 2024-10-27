package secret

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
)

// GenerateSSHKeyPair generates an RSA key pair (private and public). The private key is returned in PEM format, and the public key is returned in OpenSSH format.
func GenerateSSHKeyPair(ctx context.Context) (privateKey, publicKey string, err error) {
	// Generate RSA key
	privateKeyObj, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	// Encode private key as PEM
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKeyObj),
	})

	// Generate public key in OpenSSH format
	publicKeyObj, err := ssh.NewPublicKey(&privateKeyObj.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKeySSH := string(ssh.MarshalAuthorizedKey(publicKeyObj))

	return string(privateKeyPEM), publicKeySSH, nil
}

// GenerateRSAKeyPair generates an RSA key pair and returns the private and public keys in PEM format
func GenerateRSAKeyPair(ctx context.Context) (string, string, error) {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	// Convert private key to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Extract the public key from the private key
	publicKey := &privateKey.PublicKey

	// Convert public key to PKIX ASN.1 DER encoded form
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}

	// Convert public key to PEM format
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

func ParsePEMPrivateKey(ctx context.Context, pemPrivateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemPrivateKey))
	if block == nil {
		return nil, fmt.Errorf("Failed to decode private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParsePEMPublicKey(ctx context.Context, pemPublicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemPublicKey))
	if block == nil {
		return nil, fmt.Errorf("Failed to decode public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Failed to parse public key")
	}

	return rsaPub, nil
}

func ParseOpenSSHPublicKey(ctx context.Context, encodedKey string) (*rsa.PublicKey, error) {
	// Parse the SSH public key
	sshPubKey, err := ssh.ParsePublicKey([]byte(encodedKey))
	if err != nil {
		return nil, err
	}

	// To get back to an *rsa.PublicKey, we need to first upgrade to the
	// ssh.CryptoPublicKey interface
	parsedCryptoKey, ok := sshPubKey.(ssh.CryptoPublicKey)
	if !ok {
		return nil, fmt.Errorf("Failed to convert SSH public key to CryptoPublicKey")
	}

	// Then, we can call CryptoPublicKey() to get the actual crypto.PublicKey
	pubCrypto := parsedCryptoKey.CryptoPublicKey()

	// Finally, we can convert back to an *rsa.PublicKey
	pubKey, ok := pubCrypto.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Failed to convert CryptoPublicKey to RSA public key")
	}

	return pubKey, nil
}

// GenerateSalt generates a cryptographically secure random salt of the given length
func GenerateSalt(ctx context.Context, length int) (string, error) {
	// Create a byte slice to hold the random bytes
	salt := make([]byte, length)

	// Read random bytes into the slice
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	// Encode the salt as a base64 string to make it easier to store or transmit
	return base64.StdEncoding.EncodeToString(salt), nil
}
