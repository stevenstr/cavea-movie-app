package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrNotFound = errors.New("service address not found")

// Registry represents service registry.
type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceID, seriviceName string) error
	ServiceAddress(ctx context.Context, instanceID string) ([]string, error)
	ReportHealthyState(instanceID, seriviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
