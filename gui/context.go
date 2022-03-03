package gui

import (
	"errors"

	"github.com/prgres/ecsctl/cluster"
	"github.com/prgres/ecsctl/service"
)

type Context struct {
	ClustersData []*cluster.ClusterData

	ActiveCluster *cluster.ClusterData
	ActiveService *service.ServiceData

	Views       []*View
	CurrentView *View
}

func NewContext() (*Context, error) {
	clustersData, err := cluster.GetClusters()
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

func (ctx *Context) View(id string) (*View, error) {
	for i := range ctx.Views {
		if id == ctx.Views[i].Id {
			return ctx.Views[i], nil
		}
	}

	return nil, errors.New("view: " + id + " not found")
}

func (ctx *Context) SetCurrentView(id string) (*View, error) {
	view, err := ctx.View(id)
	if err != nil {
		return nil, err
	}
	ctx.CurrentView = view

	return ctx.CurrentView, nil
}
