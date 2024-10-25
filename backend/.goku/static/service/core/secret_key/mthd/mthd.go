package mthd

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"ongoku/backend/src/lib/secret"
	"os"
	"path/filepath"

	"github.com/teejays/gokutil/env/envutil"
	"github.com/teejays/gokutil/errutil"
	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
)

type EntityType = svccore_entsecretkey_typ.SecretKey

func HookCreatePre(ctx context.Context, c app_client.Client, in EntityType) (EntityType, error) {

	// This generates an RSA key
	if in.Type == svccore_entsecretkey_typ.Type_RSA {

		pvtKey, pubKey, err := secret.GenerateRSAKeyPair(ctx)
		if err != nil {
			return in, fmt.Errorf("Generating a new SSH Key Pair: %w", err)
		}
		in.PublicKey = pubKey
		in.Type = svccore_entsecretkey_typ.Type_RSA
		in.PrivateKeyFormat = svccore_entsecretkey_typ.Format_PEM
		in.PublicKeyFormat = svccore_entsecretkey_typ.Format_PEM

		// Save the private key as file in secrets folder
		scrtFilePath, err := GetSecretFilePath(ctx, in.ID.String())
		if err != nil {
			return in, fmt.Errorf("Getting secret file path: %w", err)
		}
		// Todo
		// scrtFilePath := path + ".pem"
		log.Debug(ctx, "[Core > SecretKey] HookCreatePre: Writing private key to file", "path", scrtFilePath)

		err = os.WriteFile(scrtFilePath, []byte(pvtKey), 0644)
		if err != nil {
			return in, fmt.Errorf("Writing private key to file: %w", err)
		}
	} else {
		return in, errutil.NewGerror(http.StatusBadRequest, "Currently only RSA type keys are supported.", "Unsupported key type: %s", in.Type)
	}

	return in, nil

}

func GetPrivateKey(ctx context.Context, secretKey svccore_entsecretkey_typ.SecretKey) (*rsa.PrivateKey, error) {
	// Get the private key string
	pvtKeyStr, err := GetPrivateKeyString(ctx, secretKey)
	if err != nil {
		return nil, fmt.Errorf("Getting private key string: %w", err)
	}
	// Parse the private key
	pvtKey, err := secret.ParsePEMPrivateKey(ctx, pvtKeyStr)
	if err != nil {
		return nil, fmt.Errorf("Parsing private key: %w", err)
	}
	return pvtKey, nil
}

func GetPrivateKeyString(ctx context.Context, secretKey svccore_entsecretkey_typ.SecretKey) (string, error) {
	scrtFilePath, err := GetSecretFilePath(ctx, secretKey.ID.String())
	if err != nil {
		return "", fmt.Errorf("Getting secret file path: %w", err)
	}
	pvtKey, err := os.ReadFile(scrtFilePath)
	if err != nil {
		return "", fmt.Errorf("Reading private key from file: %w", err)
	}
	return string(pvtKey), nil
}

func GetSecretFilePath(ctx context.Context, keyID string) (string, error) {
	appRootDir := envutil.GetEnvVarStr("APP_ROOT_DIR")
	if appRootDir == "" {
		return "", fmt.Errorf("APP_ROOT_DIR env var is not set")
	}
	scrtDirPath := envutil.GetEnvVarStr("SECRETS_DIR_PATH")
	if scrtDirPath == "" {
		return "", fmt.Errorf("SECRETS_DIR_PATH not set")
	}
	scrtFilePath := filepath.Join(appRootDir, scrtDirPath, fmt.Sprintf("secret-key-%s", keyID))
	return scrtFilePath, nil
}
