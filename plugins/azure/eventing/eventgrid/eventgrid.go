package eventgrid_service

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/eventgrid"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/eventgrid/eventgridapi"
	evtmgmt "github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/mgmt/eventgrid"
	egmgmtapi "github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/mgmt/eventgrid/eventgridapi"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/nitric-dev/membrane/plugins/azure/config"
	"github.com/nitric-dev/membrane/plugins/sdk"
	"google.golang.org/api/iterator"
)

type EventGridService struct {
	client       eventgridapi.BaseClientAPI
	topicsClient egmgmtapi.TopicsClientAPI
}

func (e *EventGridService) getTopicByName(name string) (*evtmgmt.Topic, error) {
	ctx := context.Background()

	// Use Filter string
	if res, err := e.topicsClient.ListBySubscriptionComplete(ctx, "", nil); err != nil {
		return nil, err
	} else {
		// Build list of topics then break and return
		for {
			if err := res.Next(); err != nil {
				if err == iterator.Done {
					break
				} else {
					return nil, err
				}
			}

			topic := res.Value()

			if *topic.Name == name {
				return &topic, nil
			}
		}
	}

	return nil, fmt.Errorf("Unable to find topic with name %s", name)
}

// ListTopics - Topics that belong in the configured resource group
func (e *EventGridService) ListTopics() ([]string, error) {
	ctx := context.Background()
	topics := make([]string, 0)

	if res, err := e.topicsClient.ListBySubscriptionComplete(ctx, "", nil); err != nil {
		return nil, err
	} else {
		// Build list of topics then break and return
		for {
			if err := res.Next(); err != nil {
				if err == iterator.Done {
					break
				} else {
					return nil, err
				}
			}
			topic := res.Value()
			topics = append(topics, *topic.Name)
		}
	}

	return topics, nil
}

func (e *EventGridService) Publish(topic string, evt *sdk.NitricEvent) error {
	// ctx := context.Background()

	if _, err := e.getTopicByName(topic); err == nil {
		// TODO: Create events...
		//e.client.PublishCustomEventEvents(ctx, )(ctx, *topic.Endpoint, []eventgrid.Event{
		//	eventgrid.Event{

		//	}
		//})
		return nil
	} else {
		return err
	}
}

// New - Creates a new instance of the Nitric Azure Event Grid eventing service
func New() (sdk.EventService, error) {
	authorizer, _ := auth.NewAuthorizerFromEnvironment()
	// Load Azure config from environment variables
	config := config.FromEnv()

	topicsClient := evtmgmt.NewTopicsClient(config.SubscriptionID())
	topicsClient.Authorizer = authorizer

	return &EventGridService{
		client:       eventgrid.New(),
		topicsClient: topicsClient,
	}, nil
}

func NewWithClients(client eventgridapi.BaseClientAPI, mgmt egmgmtapi.TopicsClientAPI) sdk.EventService {
	return &EventGridService{
		client:       client,
		topicsClient: mgmt,
	}
}
