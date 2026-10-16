package platform

import "context"

// Shell is the desktop host entry surface for a given OS.
type Shell interface {
	// OS returns the host platform identifier (darwin, linux, or windows).
	OS() string
	// Run starts the desktop shell. Stubs return a not-implemented error.
	Run(ctx context.Context) error
}
