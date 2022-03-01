package context

import (
	"github.com/prgres/ecsctl/cluster"
)

type Context struct {
	ClustersData []*cluster.ClusterData

	IsServiceFetched bool
	ActiveCluster    *cluster.ClusterData
}

func New() (*Context, error) {
	clustersData, err := GetClusters()
	if err != nil {
		return nil, err
	}

	return &Context{
		ClustersData: clustersData,
	}, nil

}

func (ctx *Context) Cluster(id string) (*cluster.ClusterData, error) {
	for _, cluster := range ctx.ClustersData {
		if cluster.Name == id {
			return cluster, nil
		}
	}

	return nil, cluster.ErrClusterNotFound
}

func (ctx *Context) SetActiveCluster(cluster *cluster.ClusterData) {
	ctx.ActiveCluster = cluster
}

func (ctx *Context) SetActiveClusterId(clusterId string) error {
	cluster, err := ctx.Cluster(clusterId)
	if err != nil {
		return err
	}

	ctx.SetActiveCluster(cluster)
	return nil
}

func (ctx *Context) Context() *Context {
	return ctx
}

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
