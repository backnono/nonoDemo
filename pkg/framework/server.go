package framework

import (
	"context"
)

// Server interface is the interface that must be implemented by a server.
type Server interface {
	Serve(ctx context.Context) error
}
