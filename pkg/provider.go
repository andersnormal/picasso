package pkg

import (
	"context"
)

// Provider ...
type Provider interface {
	// CloneWithContext ...
	CloneWithContext(context.Context) error
}
