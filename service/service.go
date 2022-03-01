package service

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type ServiceData struct {
	Name string
	Arn  string

	*types.Service
}

func GetServices(cluster string) ([]string, error) {
	var nextToken *string
	var servicesArn []string

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := ecs.NewFromConfig(cfg)

	for {
		output, err := svc.ListServices(ctx, &ecs.ListServicesInput{
			NextToken:  nextToken,
			MaxResults: aws.Int32(100),
			Cluster:    &cluster,
		})
		if err != nil {
			return nil, err
		}

		servicesArn = append(servicesArn, output.ServiceArns...)
		nextToken = output.NextToken
		if nextToken == nil {
			break
		}
	}

	return servicesArn, nil
}

func GetNameFromResourceId(resourceId string) string {
	parts := strings.Split(resourceId, "/")
	return parts[len(parts)-1]
}

func DescribeServices(cluster string, services []string) ([]*types.Service, error) {
	result := make([]*types.Service, len(services))
	globIndex := 0
	chunkSize := 10

	if len(services) < 10 {
		chunkSize = len(services)
	}

	for i := 0; i < len(services); i += chunkSize {
		end := i + chunkSize
		if end > len(services) {
			end = len(services)
		}

		output, err := describeService(cluster, services[i:end])
		if err != nil {
			return nil, err
		}

		for j := range output {
			result[globIndex] = &output[j]
			globIndex++
		}
	}

	return result, nil
}

func describeService(cluster string, services []string) ([]types.Service, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := ecs.NewFromConfig(cfg)

	output, err := svc.DescribeServices(ctx, &ecs.DescribeServicesInput{
		Cluster:  &cluster,
		Services: services,
	})

	if err != nil {
		return nil, err
	}

	return output.Services, nil
}
