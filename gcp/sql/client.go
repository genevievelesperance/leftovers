package sql

import (
	"fmt"

	"google.golang.org/api/googleapi"
	gcpsql "google.golang.org/api/sqladmin/v1beta4"
)

type client struct {
	project string
	logger  logger

	service   *gcpsql.Service
	instances *gcpsql.InstancesService
}

func NewClient(project string, service *gcpsql.Service, logger logger) client {
	return client{
		project:   project,
		logger:    logger,
		service:   service,
		instances: service.Instances,
	}
}

func (c client) ListInstances() (*gcpsql.InstancesListResponse, error) {
	return c.instances.List(c.project).Do()
}

func (c client) DeleteInstance(instance string) error {
	_, err := c.instances.Delete(c.project, instance).Do()
	return err
}

type request interface {
	Do(...googleapi.CallOption) (*gcpsql.Operation, error)
}

func (c client) wait(request request) error {
	op, err := request.Do()
	if err != nil {
		return fmt.Errorf("Do request: %s", err)
	}

	waiter := NewOperationWaiter(op, c.service, c.project, c.logger)

	return waiter.Wait()
}