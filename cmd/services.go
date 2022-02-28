package cmd

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
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
