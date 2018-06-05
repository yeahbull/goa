// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars service
//
// Command:
// $ goa gen goa.design/goa/examples/cars/design -o
// $(GOPATH)/src/goa.design/goa/examples/cars

package carssvc

import (
	"context"
)

// The cars service lists car models by body style.
type Service interface {
	// Lists car models by body style.
	List(context.Context, *ListPayload, ListServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "cars"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"list"}

// ListServerStream is the interface a list endpoint server stream must satisfy.
type ListServerStream interface {
	// Send sends Car result across the stream.
	Send(*Car) error
	// Close closes the result stream.
	Close() error
}

// ListClientStream is the interface a list endpoint client stream must satisfy.
type ListClientStream interface {
	// Recv receives Car result from the stream.
	Recv() (*Car, error)
}

// ListPayload is the payload type of the cars service list method.
type ListPayload struct {
	// The car body style.
	Style string
}

// Car is the result type of the cars service list method.
type Car struct {
	// The make of the car
	Make string
	// The car model
	Model string
	// The car body style
	BodyStyle string
}
