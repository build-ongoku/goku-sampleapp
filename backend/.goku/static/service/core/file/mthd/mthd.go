package mthd

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	client_file "github.com/teejays/gokutil/client/file"
	"github.com/teejays/gokutil/env/envutil"
	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/scalars"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
)

var storageClientToFileClientType = map[svccore_entfile_typ.StorageClient]client_file.ClientType{
	svccore_entfile_typ.StorageClient_LocalFile: client_file.TypeLocal,
}

// HookCreatePre is a hook that runs before adding a file
func HookCreatePre(ctx context.Context, c app_client.Client, req svccore_entfile_typ.File) (svccore_entfile_typ.File, error) {
	log.Warn(ctx, "HookCreatePre: Adding file", "file", req)
	var err error

	// Fetch the file from the storage client and inspect it to populate file metadata
	if req.StorageClient == svccore_entfile_typ.StorageClient_INVALID {
		return req, fmt.Errorf("Invalid StorageClient type")
	}

	fileClientType, ok := storageClientToFileClientType[req.StorageClient]
	if !ok {
		return req, fmt.Errorf("Unsupported StorageClient type: %s", req.StorageClient)
	}

	// Get the file client
	var fileClient client_file.IFileClient

	switch fileClientType {
	case client_file.TypeLocal:
		appRootDirPath := envutil.GetEnvVarStr("APP_ROOT_DIR")
		if appRootDirPath == "" {
			return req, fmt.Errorf("env var APP_ROOT_DIR is not set")
		}
		fileClient, err = client_file.NewLocalFileClient(ctx, filepath.Join(appRootDirPath, ".dev"))
		if err != nil {
			return req, fmt.Errorf("constructing a new file client: %w", err)
		}

	default:
		return req, fmt.Errorf("Unsupported file client type: %s", fileClientType)
	}

	fileStorageID, err := scalars.NewIDFromString(req.StorageClientIdentifier)
	if err != nil {
		return req, fmt.Errorf("parsing file ID: %w", err)
	}
	fileRdr, err := fileClient.FileRead(ctx, fileStorageID)
	if err != nil {
		return req, fmt.Errorf("reading file: %w", err)
	}
	defer fileRdr.Close()

	fileData, err := io.ReadAll(fileRdr)
	if err != nil {
		return req, fmt.Errorf("reading file data: %w", err)
	}

	// Update the file metadata
	// Size
	req.SizeBytes = len(fileData)
	// Hash
	hash := sha256.New()
	req.FileHash = fmt.Sprintf("%x", hash.Sum(fileData))
	// Mime type
	req.MimeType = http.DetectContentType(fileData)

	return req, nil
}
