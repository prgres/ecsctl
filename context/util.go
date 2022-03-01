package context

import "github.com/prgres/ecsctl/cluster"

func GetClusters() ([]*cluster.ClusterData, error) {
	var clustersData []*cluster.ClusterData

	clusters, err := cluster.GetClustersNames()
	if err != nil {
		return nil, err
	}

	for _, c := range clusters {
		clustersData = append(clustersData, &cluster.ClusterData{
			Name: c,
		})
	}

	return clustersData, nil
}
