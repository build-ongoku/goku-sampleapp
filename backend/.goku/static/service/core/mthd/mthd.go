package svccore_mthd

import (
	"context"
	"fmt"
	"net/http"

	client_file "github.com/teejays/gokutil/client/file"
	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
)

func FileUpload(ctx context.Context, c app_client.Client, req *http.Request) (svccore_entfile_typ.File, error) {

	var resp svccore_entfile_typ.File
	var err error

	// Get the file from the request
	file, fileRequestHandler, err := req.FormFile("file")
	if err != nil {
		return resp, fmt.Errorf("FileUpload: getting file from request: %w", err)
	}

	log.Debug(ctx, "FileUpload: Got the file from the request", "file", file, "fileInfo", fileRequestHandler)

	fileClient, err := client_file.GetDefaultFileClient(ctx)
	if err != nil {
		return resp, fmt.Errorf("Getting default file client: %w", err)
	}

	fileID, err := fileClient.FileUpload(ctx, file)
	if err != nil {
		return resp, fmt.Errorf("FileUpload: uploading file: %w", err)
	}

	log.Debug(ctx, "FileUpload: File uploaded successfully", "fileID", fileID)

	// Mapping of FileClientType to StorageClient

	storageClientType, ok := fileClientTypeToStorageClient[fileClient.GetClientType()]
	if !ok {
		return resp, fmt.Errorf("FileUpload: Unsupported file client type: %s", fileClient.GetClientType())
	}

	// Should we create a file entity?
	fileReq := svccore_entfile_typ.AddRequest{
		Object: svccore_entfile_typ.FileInput{
			FileName:                fileRequestHandler.Filename,
			StorageClient:           storageClientType,
			StorageClientIdentifier: fileID.String(),
		},
	}

	resp, err = c.Core().File().Add(ctx, fileReq)
	if err != nil {
		return resp, fmt.Errorf("creating file entity: %w", err)
	}

	return resp, nil
}

func SecretDecryptabeAdd(ctx context.Context, c app_client.Client, req svccore_typ.SecretDecryptableAddRequest) (svccore_entsecretdecryptable_typ.SecretDecryptable, error) {
	var ret svccore_entsecretdecryptable_typ.SecretDecryptable
	var err error

	return ret, err
}

var fileClientTypeToStorageClient = map[client_file.ClientType]svccore_entfile_typ.StorageClient{
	client_file.TypeLocal: svccore_entfile_typ.StorageClient_LocalFile,
}
