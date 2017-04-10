Golang dependency injection
------
A dependency injection container implementation.
Based on [PHP PSR specification](https://github.com/container-interop/fig-standards/blob/master/proposed/container.md)

### Features
- [x] Get/Has methods
- [x] Delegate Container

### Usage
The container stores types that implement the `ServiceFactoryInterface` which are responsible
to return the instance of the service.

A service can implement himself the `ServiceFactoryInterface` so that the instance is shared between calls.

When a delegate container is defined the depency injection will be delegated to the delegate container.

Since the `Get` method returns an `interface{}` type assertion should be used to resolve the real type
of the Service.

#### Code example

Create a route and then fetch the matched route with the parameter
``` go
container := di.NewContainer()

container.AddService("serviceName", &SomeServiceFactory{})

if container.Has("serviceName") {
    service := container.Get("serviceName")
}
```
