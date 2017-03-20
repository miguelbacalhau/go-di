// Package di provides a simple service container implementation
// to be used for dependency injection in external packages.
package di

// Base service interface serves only has a representation.
type ServiceInterface interface {
}

// Standard Container interface
type ContainerInterface interface {
	Get(service string) interface{}
	Has(service string) bool
}

// Service Factory responsible for the Service instantiation.
// All Service Factories must implement this interface
type ServiceFactoryInterface interface {
	CreateService(container ContainerInterface) interface{}
}

// Simple implementation of the container interface using
// a map of string to Service Factories
type Container struct {
	services          map[string]ServiceFactoryInterface
	delegateContainer ContainerInterface
}

// Container Get Method implementation, returns the instance
// of the Service registered in the Container with the name
// provided.
func (container *Container) Get(name string) interface{} {
	var dependencyContainer ContainerInterface

	if container.delegateContainer != nil {
		dependencyContainer = container.delegateContainer
	} else {
		dependencyContainer = container
	}
	// TODO not found error instead of failing
	return container.services[name].CreateService(dependencyContainer)
}

// Container Has Method implementation, verifies if the Service
// is registered in the Container with the name
// provided
func (container *Container) Has(name string) bool {
	_, found := container.services[name]

	return found
}

// Container Add method adds a new service entry mapped
// to the given name
func (container *Container) AddService(name string, service ServiceFactoryInterface) {
	container.services[name] = service
}

// Container constructor, creates a new Container instance
// if not needed then the delegate container should nil
func NewContainer(services map[string]ServiceFactoryInterface, delegateContainer ContainerInterface) *Container {
	return &Container{
		services:          services,
		delegateContainer: delegateContainer,
	}
}
