package awslambda

import (
	"github.com/nitrictech/nitric/engines/terraform/plugins/service"
)

type AwsLambdaService struct {
}

var _ service.Service = (*AwsLambdaService)(nil)

func (s *AwsLambdaService) Start(proxy service.Proxy) error {
	return nil
}

func New() *AwsLambdaService {
	return &AwsLambdaService{}
}
