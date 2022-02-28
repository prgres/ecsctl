package cmd

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

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
	var result []*types.Service

	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := ecs.NewFromConfig(cfg)

	chunks := 10
	if len(services) > 10 {
		chunks = len(services) / 10
		if len(services)%10 != 0 {
			chunks++
		}
	}

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
			// []string{
			// "eden-kafka-schema-registry-dev-ServiceDefinitionInternalELB-1QY0GMMC0OX3Q",
			// },
			//  services[:10], //TODO: can be max 10. split into 10 parts chunks
		}
		output, err := svc.DescribeServices(ctx, &tmp)
		if err != nil {
			return nil, err
		}

		for _, r := range output.Services {
			result = append(result, &r)
		}
	}

	// offset := 10
	// bufferLen := func() int {
	// 	if len(cluster) <= offset {
	// 		return 1
	// 	}
	// 	t := len(cluster) / offset
	// 	if len(cluster)%offset > 0 {
	// 		t++
	// 	}
	// 	return t
	// }()

	// for index := 0; index < bufferLen; index++ {
	// 	output, err := svc.DescribeServices(ctx, &ecs.DescribeServicesInput{
	// 		Cluster:  &cluster,
	// 		Services: services[index*offset : index*offset+offset],
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for _, r := range output.Services {
	// 		result = append(result, &r)
	// 	}
	// 	break
	// }

	return result, nil
}
