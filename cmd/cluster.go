package cmd

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func GetClusters() ([]string, error) {
	var nextToken *string
	var clustersArn []string

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := ecs.NewFromConfig(cfg)

	for {
		output, err := svc.ListClusters(ctx, &ecs.ListClustersInput{
			NextToken: nextToken,
		})
		if err != nil {
			return nil, err
		}

		clustersArn = append(clustersArn, output.ClusterArns...)
		nextToken = output.NextToken
		if nextToken == nil {
			break
		}
	}

	return clustersArn, nil
}

func ParseClusterArn(arn string) string {
	re := regexp.MustCompile(`arn:.*:cluster/`)
	return re.ReplaceAllString(arn, "")
}

func GetClustersNames() ([]string, error) {
	clusters, err := GetClusters()
	if err != nil {
		return nil, err
	}

	for i, cluster := range clusters {
		clusters[i] = ParseClusterArn(cluster)
	}

	return clusters, nil
}
