package blobstorage_service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/nitric-dev/membrane/plugins/sdk"
	"github.com/nitric-dev/membrane/utils"
)

// BlobStorageService - Is the concrete implementation of Azure Blob Storage for the Nitric Storage Plugin
type BlobStorageService struct {
	sdk.UnimplementedStoragePlugin
	client azblob.ServiceURL
}

// Put - Writes a new item to a bucket
func (s *BlobStorageService) Put(bucket string, key string, object []byte) error {
	ctx := context.Background()

	containerURL := s.client.NewContainerURL(bucket)
	blobURL := containerURL.NewBlockBlobURL(key)

	if _, err := azblob.UploadBufferToBlockBlob(ctx, object, blobURL, azblob.UploadToBlockBlobOptions{}); err != nil {
		return err
	}

	return nil
}

// Get - Retrieves an item from a bucket
func (s *BlobStorageService) Get(bucket string, key string) ([]byte, error) {
	ctx := context.Background()

	containerURL := s.client.NewContainerURL(bucket)
	blobURL := containerURL.NewBlobURL(key)

	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})

	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	bytes, err := ioutil.ReadAll(bodyStream)

	return bytes, err
}

// Delete - Deletes an item from a bucket
func (s *BlobStorageService) Delete(bucket string, key string) error {
	ctx := context.Background()

	containerURL := s.client.NewContainerURL(bucket)
	blobURL := containerURL.NewBlockBlobURL(key)

	_, err := blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionType(""), azblob.BlobAccessConditions{})

	return err
}

// New creates a new default S3 storage plugin
func New() (sdk.StorageService, error) {
	// TODO: Create a default storage account for the stack???
	storageAccount := utils.GetEnv("AZURE_STORAGE_ACCOUNT", "")
	storageKey := utils.GetEnv("AZURE_STORAGE_ACCESS_KEY", "")

	if len(storageAccount) == 0 || len(storageKey) == 0 {
		return nil, fmt.Errorf("Could not initialise Azure Blob Storage service, missing AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY variables")
	}

	var credential *azblob.SharedKeyCredential
	var accountURL *url.URL
	var err error

	if credential, err = azblob.NewSharedKeyCredential(storageAccount, storageKey); err != nil {
		return nil, err
	}

	if accountURL, err = url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", storageAccount)); err != nil {
		return nil, err
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	client := azblob.NewServiceURL(*accountURL, pipeline)

	return &BlobStorageService{
		client: client,
	}, nil
}
