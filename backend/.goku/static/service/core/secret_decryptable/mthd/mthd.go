package mthd

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/scalars"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_entsecretdecryptable_typhelper "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typhelper"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
	svccore_entsecretkey_typhelper "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typhelper"
	secretlib "sampleapp/backend/.goku/static/lib/secret"
	secretkeymthd "sampleapp/backend/.goku/static/service/core/secret_key/mthd"
)

type EntityType = svccore_entsecretdecryptable_typ.SecretDecryptable

func HookCreatePre(ctx context.Context, c app_client.Client, in EntityType) (EntityType, error) {

	// Salt the secret
	salt, err := generateSalt(ctx)
	if err != nil {
		return in, fmt.Errorf("generating salt: %w", err)
	}
	saltedRawValue := saltValue(ctx, in.RawValue, salt)
	in.Salt = salt

	// Get the secret key for encryption
	secretKey, err := GetSecretKeyByIDOrRandom(ctx, c, in.SecretKeyID)
	if err != nil {
		return in, fmt.Errorf("getting secret key: %w", err)
	}

	// Encrypt the secret
	encryptedValue, err := EncryptValueUsingSecretKey(ctx, saltedRawValue, secretKey)
	if err != nil {
		return in, fmt.Errorf("encrypting secret data: %w", err)
	}

	in.SecretKeyID = secretKey.ID
	in.EncryptedValue = encryptedValue

	return in, nil
}

func ActionDecrypt(ctx context.Context, c app_client.Client, in svccore_entsecretdecryptable_typ.StandardEntityRequest) (svccore_entsecretdecryptable_typ.StandardEntityResponse, error) {
	var ret svccore_entsecretdecryptable_typ.StandardEntityResponse

	// Get the secret
	secret, err := svccore_entsecretdecryptable_typhelper.GetByID(ctx, c.Core().SecretDecryptable(), in.ID)
	if err != nil {
		return ret, fmt.Errorf("getting secret: %w", err)
	}

	// Get the secret key
	secretKey, err := svccore_entsecretdecryptable_typhelper.GetSecretKey(ctx, c.Core().SecretKey(), secret)
	if err != nil {
		return ret, fmt.Errorf("getting secret key: %w", err)
	}

	// Decrypt the secret

	decryptedValue, err := DecryptValueUsingSecretKey(ctx, secret.EncryptedValue, secretKey)
	if err != nil {
		return ret, fmt.Errorf("decrypting secret data: %w", err)
	}

	// Desalt the secret
	rawValue := desaltValue(ctx, decryptedValue, secret.Salt)

	secret.RawValue = rawValue

	ret.Object = secret

	return ret, nil
}

func GetSecretKeyByIDOrRandom(ctx context.Context, c app_client.Client, id scalars.ID) (svccore_entsecretkey_typ.SecretKey, error) {
	var err error

	// Choose a Secret Key for encryption. Either use the one provided in the request or choose a random one.
	var secretKey svccore_entsecretkey_typ.SecretKey
	if !id.IsEmpty() {
		secretKey, err = svccore_entsecretkey_typhelper.GetByID(ctx, c.Core().SecretKey(), id)
		if err != nil {
			return secretKey, fmt.Errorf("getting secret key: %w", err)
		}
	} else {
		// choose a (RSA) secret key
		listResp, err := c.Core().SecretKey().List(ctx, svccore_entsecretkey_typ.ListRequest{
			Filter: svccore_entsecretkey_typ.SecretKeyFilter{
				Type: &svccore_entsecretkey_typ.TypeCondition{
					Op: filter.EQUAL,
					Values: []svccore_entsecretkey_typ.Type{
						svccore_entsecretkey_typ.Type_RSA,
					},
				},
			},
		})
		if err != nil {
			return secretKey, fmt.Errorf("getting secret key: %w", err)
		}
		if len(listResp.Items) == 0 {
			return secretKey, fmt.Errorf("No secret key found in the system. Please specify one in the request.")
		}
		secretKey = listResp.Items[0]
		log.Debug(ctx, "Using a random secret key", "key", secretKey)
	}

	return secretKey, nil

}

func EncryptValueUsingSecretKey(ctx context.Context, valueToEncrypt string, secretKey svccore_entsecretkey_typ.SecretKey) (string, error) {
	var err error

	// Encrypt the secret data using the secret key
	if secretKey.PublicKey == "" {
		return "", fmt.Errorf("Secret key does not have a public key")
	}

	// Parse the public key
	var parsedPublicKey *rsa.PublicKey
	switch secretKey.PublicKeyFormat {
	case svccore_entsecretkey_typ.Format_PEM:
		parsedPublicKey, err = secretlib.ParsePEMPublicKey(ctx, secretKey.PublicKey)
		if err != nil {
			return "", fmt.Errorf("parsing public key: %w", err)
		}
	case svccore_entsecretkey_typ.Format_OpenSSH:
		parsedPublicKey, err = secretlib.ParseOpenSSHPublicKey(ctx, secretKey.PublicKey)
		if err != nil {
			return "", fmt.Errorf("parsing public key: %w", err)
		}
	default:
		return "", fmt.Errorf("Unsupported public key format: %s", secretKey.PublicKeyFormat)
	}

	// Encrypt the data with the public key
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, parsedPublicKey, []byte(valueToEncrypt))
	if err != nil {
		return "", fmt.Errorf("encrypting data: %w", err)
	}

	// Return the encrypted data as a base64 encoded string
	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encryptedBase64, nil

}

func DecryptValueUsingSecretKey(ctx context.Context, encryptedValue string, secretKey svccore_entsecretkey_typ.SecretKey) (string, error) {
	var err error

	// Get the parsed private key
	pvtKey, err := secretkeymthd.GetPrivateKey(ctx, secretKey)
	if err != nil {
		return "", fmt.Errorf("getting private key: %w", err)
	}

	// Decode the base64 encoded string
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedValue)
	if err != nil {
		return "", fmt.Errorf("decoding base64: %w", err)
	}

	// Decrypt the data with the private key
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, pvtKey, encryptedBytes)
	if err != nil {
		return "", fmt.Errorf("decrypting data: %w", err)
	}

	// Return the decrypted data as a string
	return string(decryptedBytes), nil

}

func generateSalt(ctx context.Context) (string, error) {
	return secretlib.GenerateSalt(ctx, 32)
}

func saltValue(ctx context.Context, rawValue, salt string) string {
	return rawValue + "---" + salt
}

func desaltValue(ctx context.Context, saltedValue, salt string) string {
	return strings.TrimSuffix(saltedValue, "---"+salt)
}
