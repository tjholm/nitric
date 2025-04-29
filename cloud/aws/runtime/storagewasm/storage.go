package storagewasm

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	sdk "github.com/extism/go-sdk"
	"google.golang.org/grpc/codes"

	"github.com/nitrictech/nitric/cloud/aws/runtime/resource"
	grpc_errors "github.com/nitrictech/nitric/core/pkg/grpc/errors"
	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
)

//go:embed plugin/dist/plugin.wasm
var wasmPluginBytes []byte

// StorageWasmService implements the StorageServer interface using WASM plugins
type StorageWasmService struct {
	plugin   *sdk.Plugin
	resolver resource.AwsResourceResolver
}

var _ storagepb.StorageServer = &StorageWasmService{}

func (s *StorageWasmService) getS3BucketName(ctx context.Context, bucket string) (*string, error) {
	fmt.Println("getS3BucketName", bucket)
	buckets, err := s.resolver.GetResources(ctx, resource.AwsResource_Bucket)
	if err != nil {
		return nil, fmt.Errorf("error getting bucket list: %w", err)
	}

	if s3Bucket, ok := buckets[bucket]; ok {
		bucketName := strings.Split(s3Bucket.ARN, ":::")[1]

		return aws.String(bucketName), nil
	}

	return nil, fmt.Errorf("bucket %s does not exist", bucket)
}

// New creates a new StorageWasmService
func New(resolver resource.AwsResourceResolver) (*StorageWasmService, error) {
	// Load the WASM plugin
	manifest := sdk.Manifest{
		Wasm: []sdk.Wasm{
			sdk.WasmData{
				Data: wasmPluginBytes,
			},
		},
	}

	plugin, err := sdk.NewPlugin(context.Background(), manifest, sdk.PluginConfig{
		EnableWasi: true,
	}, nil)
	if err != nil {
		fmt.Println("failed to create WASM plugin", err)
	}

	return &StorageWasmService{
		plugin:   plugin,
		resolver: resolver,
	}, nil
}

// Read implements the StorageServer interface
func (s *StorageWasmService) Read(ctx context.Context, req *storagepb.StorageReadRequest) (*storagepb.StorageReadResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.Read")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName

	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	exit, res, err := s.plugin.Call("read", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	// Unmarshal response
	var resp storagepb.StorageReadResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}

// Write implements the StorageServer interface
func (s *StorageWasmService) Write(ctx context.Context, req *storagepb.StorageWriteRequest) (*storagepb.StorageWriteResponse, error) {
	fmt.Println("writing to bucket", req.BucketName)

	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.Write")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName
	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	fmt.Println("calling write plugin", req.BucketName)

	if s.plugin == nil {
		return nil, newErr(codes.Internal, "failed to call WASM plugin (plugin is nil)", nil)
	}

	exit, res, err := s.plugin.Call("write", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	fmt.Println("called write plugin", req.BucketName)

	// Unmarshal response
	var resp storagepb.StorageWriteResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}

// Delete implements the StorageServer interface
func (s *StorageWasmService) Delete(ctx context.Context, req *storagepb.StorageDeleteRequest) (*storagepb.StorageDeleteResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.Delete")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName

	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	exit, res, err := s.plugin.Call("delete_", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	// Unmarshal response
	var resp storagepb.StorageDeleteResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}

// PreSignUrl implements the StorageServer interface
func (s *StorageWasmService) PreSignUrl(ctx context.Context, req *storagepb.StoragePreSignUrlRequest) (*storagepb.StoragePreSignUrlResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.PreSignUrl")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName

	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	exit, res, err := s.plugin.Call("preSignUrl", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	// Unmarshal response
	var resp storagepb.StoragePreSignUrlResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}

// ListBlobs implements the StorageServer interface
func (s *StorageWasmService) ListBlobs(ctx context.Context, req *storagepb.StorageListBlobsRequest) (*storagepb.StorageListBlobsResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.ListBlobs")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName

	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	exit, res, err := s.plugin.Call("listBlobs", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	// Unmarshal response
	var resp storagepb.StorageListBlobsResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}

// Exists implements the StorageServer interface
func (s *StorageWasmService) Exists(ctx context.Context, req *storagepb.StorageExistsRequest) (*storagepb.StorageExistsResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("StorageWasmService.Exists")

	bucketName, err := s.getS3BucketName(ctx, req.BucketName)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to get bucket name", err)
	}

	req.BucketName = *bucketName

	// Marshal request to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, newErr(codes.Internal, "failed to marshal request", err)
	}

	// Call the WASM plugin
	exit, res, err := s.plugin.Call("exists", reqBytes)
	if err != nil {
		return nil, newErr(codes.Internal, fmt.Sprintf("failed to call WASM plugin (exit code: %d)", exit), err)
	}

	// Unmarshal response
	var resp storagepb.StorageExistsResponse
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, newErr(codes.Internal, "failed to unmarshal response", err)
	}

	return &resp, nil
}
