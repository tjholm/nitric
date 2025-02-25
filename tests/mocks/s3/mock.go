// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock_s3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type MockBucket struct {
	Name string
	Tags map[string]string
}

// MockS3Client - Provides an S3API compliant mock interface
type MockS3Client struct {
	sync.RWMutex
	s3iface.S3API
	buckets []*MockBucket
	storage *map[string]map[string][]byte
}

func (s *MockS3Client) DeleteObject(in *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	s.RLock()
	defer s.RUnlock()
	bucketName := in.Bucket
	key := in.Key

	for _, b := range s.buckets {
		// We found the bucketName
		if *bucketName == b.Name {
			bucket := (*s.storage)[b.Name]
			// We found the object, delete it
			if _, ok := bucket[*key]; ok {
				delete(bucket, *key)
			}
			return &s3.DeleteObjectOutput{}, nil
		}
	}

	return nil, fmt.Errorf("bucket does not exist")
}

func (s *MockS3Client) ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	buckets := make([]*s3.Bucket, 0)

	for _, b := range s.buckets {
		buckets = append(buckets, &s3.Bucket{
			Name: aws.String(b.Name),
		})
	}

	return &s3.ListBucketsOutput{
		Buckets: buckets,
	}, nil
}

func (s *MockS3Client) GetBucketTagging(in *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	bucketName := in.Bucket

	for _, b := range s.buckets {
		if b.Name == *bucketName {
			tags := make([]*s3.Tag, 0)

			for key, val := range b.Tags {
				tags = append(tags, &s3.Tag{
					Key:   aws.String(key),
					Value: aws.String(val),
				})
			}

			return &s3.GetBucketTaggingOutput{
				TagSet: tags,
			}, nil
		}
	}

	return nil, fmt.Errorf("Bucket does not exist")
}

func (s *MockS3Client) PutObject(in *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	s.Lock()
	defer s.Unlock()
	bucket := in.Bucket
	key := in.Key
	reader := in.Body

	for _, b := range s.buckets {
		// We found the bucket
		if *bucket == b.Name {
			store := *s.storage
			if store[b.Name] == nil {
				store[b.Name] = make(map[string][]byte)
			}
			storeKey := *key
			bytes, _ := ioutil.ReadAll(reader)

			store[b.Name][storeKey] = bytes

			return &s3.PutObjectOutput{}, nil
		}
	}

	return nil, fmt.Errorf("bucket does not exist")
}

func (s *MockS3Client) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	s.RLock()
	defer s.RUnlock()
	bucketName := in.Bucket
	key := in.Key

	for _, b := range s.buckets {
		// We found the bucketName
		if *bucketName == b.Name {
			bucket := (*s.storage)[b.Name]
			// We found the object, by key
			if object, ok := bucket[*key]; ok {
				body := ioutil.NopCloser(bytes.NewReader(object))
				output := s3.GetObjectOutput{
					Body: body,
				}
				return &output, nil
			} else {
				return nil, fmt.Errorf("key does not exists in bucket %s", *bucketName)
			}
		}
	}

	return nil, fmt.Errorf("bucket does not exist")
}

func NewStorageClient(buckets []*MockBucket, storage *map[string]map[string][]byte) s3iface.S3API {
	return &MockS3Client{
		buckets: buckets,
		storage: storage,
	}
}
