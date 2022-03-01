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
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := ecs.NewFromConfig(cfg)

	if len(services) < 10 {
		tmp := ecs.DescribeServicesInput{
			Cluster:  &cluster,
			Services: services,
		}

		output, err := svc.DescribeServices(ctx, &tmp)
		if err != nil {
			return nil, err
		}

		for i := range output.Services {
			result[i] = &output.Services[i]
		}

		return result, nil
	}

	chunks := len(services) / 10
	if len(services)%10 != 0 {
		chunks++
	}

	globIndex := 0
	for i := 0; i < chunks; i++ {
		//TODO refactor
		rangeStart := i * 10
		rangeEnd := ((1 + i) * 10)
		var chu []string
		if i+1 < chunks {
			chu = services[rangeStart:rangeEnd]
		} else {
			chu = services[rangeStart:]
		}

		tmp := ecs.DescribeServicesInput{
			Cluster:  &cluster,
			Services: chu,
		}
		output, err := svc.DescribeServices(ctx, &tmp)
		if err != nil {
			return nil, err
		}

		for i := range output.Services {
			result[globIndex] = &output.Services[i]
			globIndex++
		}
	}

	return result, nil
}
