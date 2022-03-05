package cluster

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func GetClusters(ctx *context.Context, cfg *aws.Config) ([]*ClusterData, error) {
	var clustersData []*ClusterData

	clusters, err := getClustersNames(ctx, cfg)
	if err != nil {
		return nil, err
	}

	for _, c := range clusters {
		clustersData = append(clustersData, &ClusterData{
			Name: c,
		})
	}

	return clustersData, nil
}

func getClustersNames(ctx *context.Context, cfg *aws.Config) ([]string, error) {
	clusters, err := awsGetClusters(ctx, cfg)
	if err != nil {
		return nil, err
	}

	for i, cluster := range clusters {
		clusters[i] = parseClusterArn(cluster)
	}

	return clusters, nil
}

func parseClusterArn(arn string) string {
	re := regexp.MustCompile(`arn:.*:cluster/`)
	return re.ReplaceAllString(arn, "")
}

func awsGetClusters(ctx *context.Context, cfg *aws.Config) ([]string, error) {
	var nextToken *string
	var clustersArn []string

	svc := ecs.NewFromConfig(*cfg)

	for {
		output, err := svc.ListClusters(*ctx, &ecs.ListClustersInput{
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
