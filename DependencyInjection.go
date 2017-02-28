// Package di provides a simple service container implementation
// to be used for dependency injection in external packages.
package di

// Base service interface serves only has a representation.
type ServiceInterface interface {
}

// Standard Container interface
type ContainerInterface interface {
	Get(string) ServiceInterface
	Has(string) bool
}

// Service Factory responsible for the Service instantiation.
// All Service Factories must implement this interface
type ServiceFactoryInterface interface {
	CreateService(container ContainerInterface) ServiceInterface
}

// Simple implementation of the container interface using
// a map of string to Service Factories
type Container struct {
	Services map[string]ServiceFactoryInterface
}

// Container Get Method implementation, returns the instance
// of the Service registered in the Container with the name
// provided.
func (container *Container) Get(name string) ServiceInterface {
	// TODO not found error instead of failing
	return container.Services[name].CreateService(container)
}

// Container Has Method implementation, verifies if the Service
// is registered in the Container with the name
// provided
func (container *Container) Has(name string) bool {
	_, found := container.Services[name]

	return found
}
