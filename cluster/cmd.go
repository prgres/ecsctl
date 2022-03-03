package cluster

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func GetClusters() ([]*ClusterData, error) {
	var clustersData []*ClusterData

	clusters, err := getClustersNames()
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

func getClustersNames() ([]string, error) {
	clusters, err := awsGetClusters()
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

func awsGetClusters() ([]string, error) {
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
