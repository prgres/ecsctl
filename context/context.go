package context

import (
	"github.com/prgres/ecsctl/cluster"
	"github.com/prgres/ecsctl/service"
)

type Context struct {
	ClustersData []*cluster.ClusterData

	ActiveCluster *cluster.ClusterData
	ActiveService *service.ServiceData
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

func (ctx *Context) SetActiveCluster(cluster *cluster.ClusterData) *cluster.ClusterData {
	ctx.ActiveCluster = cluster
	return ctx.ActiveCluster
}

func (ctx *Context) SetActiveClusterId(clusterId string) (*cluster.ClusterData, error) {
	cluster, err := ctx.Cluster(clusterId)
	if err != nil {
		return nil, err
	}

	return ctx.SetActiveCluster(cluster), nil
}

func (ctx *Context) Context() *Context {
	return ctx
}
