package pkg

import (
	"context"
	"time"
)

// ProviderOpt ...
type ProviderOpt func(*ProviderOpts)

// Opts ...
type ProviderOpts struct {
	Timeout time.Duration
	URL string
}

// Provider ...
type Provider interface {
	// CloneWithContext ...
	CloneWithContext(context.Context) error
}
