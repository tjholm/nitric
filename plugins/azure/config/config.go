package config

import "github.com/nitric-dev/membrane/utils"

var (
	config *AzureConfig
)

type AzureConfig struct {
	subscriptionID string
}

func (c *AzureConfig) SubscriptionID() string {
	return c.subscriptionID
}

func FromEnv() *AzureConfig {
	if config == nil {
		subscriptionID := utils.GetEnv("AZURE_SUBSCRIPTION_ID", "")

		config = &AzureConfig{
			subscriptionID: subscriptionID,
		}
	}

	return config
}
