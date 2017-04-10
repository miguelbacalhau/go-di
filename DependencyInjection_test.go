package di

import (
	"testing"
)

type ServiceFactoryTest struct {
}

func (factory *ServiceFactoryTest) CreateService(container ContainerInterface) interface{} {
	return true
}

func TestHasService(test *testing.T) {
	services := make(map[string]ServiceFactoryInterface)
	services["test"] = &ServiceFactoryTest{}
	container := NewContainer(services, nil)

	if !container.Has("test") {
		test.Fail()
	}
}

func TestNoHasService(test *testing.T) {
	services := make(map[string]ServiceFactoryInterface)
	services["test"] = &ServiceFactoryTest{}
	container := NewContainer(services, nil)

	if container.Has("testNo") {
		test.Fail()
	}
}

func TestGetService(test *testing.T) {
	services := make(map[string]ServiceFactoryInterface)
	services["test"] = &ServiceFactoryTest{}
	container := NewContainer(services, nil)

	serviceInterface := container.Get("test")
	service := serviceInterface.(bool)

	if !service {
		test.Fail()
	}

}

func TestAddService(test *testing.T) {
	test.Fail()
}

func TestDelegateContainer(test *testing.T) {
	test.Fail()
}
