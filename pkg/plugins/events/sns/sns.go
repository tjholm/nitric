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

package sns_service

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"

	"github.com/nitrictech/nitric/pkg/plugins/errors"
	"github.com/nitrictech/nitric/pkg/plugins/errors/codes"
	"github.com/nitrictech/nitric/pkg/plugins/events"
	"github.com/nitrictech/nitric/pkg/providers/aws/core"
	utils2 "github.com/nitrictech/nitric/pkg/utils"
)

type SnsEventService struct {
	events.UnimplementedeventsPlugin
	client   snsiface.SNSAPI
	provider core.AwsProvider
}

func (s *SnsEventService) getTopics() (map[string]string, error) {
	return s.provider.GetResources(core.AwsResource_Topic)
}

// Publish to a given topic
func (s *SnsEventService) Publish(topic string, event *events.NitricEvent) error {
	newErr := errors.ErrorsWithScope(
		"SnsEventService.Publish",
		map[string]interface{}{
			"topic": topic,
			"event": event,
		},
	)

	data, err := json.Marshal(event)

	if err != nil {
		return newErr(
			codes.Internal,
			"error marshalling event payload",
			err,
		)
	}

	topics, err := s.getTopics()

	if err != nil {
		return newErr(
			codes.Internal,
			"error finding topics",
			err,
		)
	}

	topicArn, ok := topics[topic]

	if !ok {
		return newErr(
			codes.NotFound,
			"could not find topic",
			nil,
		)
	}

	message := string(data)

	publishInput := &sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message:  &message,
		// MessageStructure: json is for an AWS specific JSON format,
		// which sends different messages to different subscription types. Don't use it.
		// MessageStructure: aws.String("json"),
	}

	_, err = s.client.Publish(publishInput)

	if err != nil {
		return newErr(
			codes.Internal,
			"unable to publish message",
			err,
		)
	}

	return nil
}

func (s *SnsEventService) ListTopics() ([]string, error) {
	newErr := errors.ErrorsWithScope("SnsEventService.ListTopics", nil)

	topics, err := s.getTopics()

	if err != nil {
		return nil, newErr(
			codes.Internal,
			"error retrieving topics",
			err,
		)
	}

	topicNames := make([]string, 0, len(topics))
	for name := range topics {
		// TODO: Extract topic name from ARN
		topicNames = append(topicNames, name)
	}

	return topicNames, nil
}

// Create new SNS event service plugin
func New(provider core.AwsProvider) (events.EventService, error) {
	awsRegion := utils2.GetEnv("AWS_REGION", "us-east-1")

	sess, sessionError := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if sessionError != nil {
		return nil, fmt.Errorf("error creating new AWS session %v", sessionError)
	}

	snsClient := sns.New(sess)

	return &SnsEventService{
		client:   snsClient,
		provider: provider,
	}, nil
}

func NewWithClient(provider core.AwsProvider, client snsiface.SNSAPI) (events.EventService, error) {
	return &SnsEventService{
		provider: provider,
		client:   client,
	}, nil
}
