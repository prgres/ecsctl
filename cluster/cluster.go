package cluster

import (
	"errors"

	"github.com/prgres/ecsctl/service"
)

type ClusterData struct {
	Name     string
	Services []*service.ServiceData

	IsServiceFetched bool
}

var (
	ErrClusterNotFound = errors.New("cluster not found")
)

func (c *ClusterData) FetchServices() error {
	if c.IsServiceFetched {
		return nil
	}

	if err := c.FetchServicesNameArn(); err != nil {
		return err
	}

	if err := c.FetchServicesDetail(); err != nil {
		return err
	}

	c.IsServiceFetched = true

	return nil
}

func (c *ClusterData) FetchServicesNameArn() error {
	services, err := service.GetServices(c.Name)
	if err != nil {
		return err
	}

	for _, s := range services {
		c.Services = append(c.Services, &service.ServiceData{
			Name: service.GetNameFromResourceId(s),
			Arn:  s,
		})
	}

	return nil
}

func (c *ClusterData) FetchServicesDetail() error {
	services, err := service.DescribeServices(c.Name, c.ServicesArn())
	if err != nil {
		return err
	}

	for i, service := range c.Services {
		for _, extendedService := range services {
			if service.Arn == *extendedService.ServiceArn {
				c.Services[i].Service = extendedService
				break
			}
		}

		if c.Services[i].Service == nil {
			return errors.New("no detail found")
		}
	}

	return nil
}

func (c *ClusterData) ServicesArn() []string {
	result := make([]string, len(c.Services))

	for i, s := range c.Services {
		result[i] = s.Arn
	}

	return result
}
