//go:build darwin

package platform

import (
	"context"
	"fmt"
)

type darwinShell struct{}

// New returns the macOS (darwin) desktop shell stub.
func New() Shell {
	return darwinShell{}
}

func (darwinShell) OS() string {
	return "darwin"
}

func (darwinShell) Run(ctx context.Context) error {
	_ = ctx
	return fmt.Errorf("platform darwin (macOS): not implemented")
}
