package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewConfig() aws.Config {

	//#################### para testar local #####################
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service,
			region string, options ...interface{}) (aws.Endpoint, error) {
			if service == sqs.ServiceID {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           "http://localhost:4566",
					SigningRegion: "us-east-1",
				}, nil
			}
			return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")

		})))
	if err != nil {
		panic(err)
	}

	// production

	// cfg, err := config.LoadDefaultConfig(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	return cfg
}
