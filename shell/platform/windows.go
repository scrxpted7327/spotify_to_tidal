//go:build windows

package platform

import (
	"context"
	"fmt"
)

type windowsShell struct{}

// New returns the Windows desktop shell stub.
func New() Shell {
	return windowsShell{}
}

func (windowsShell) OS() string {
	return "windows"
}

func (windowsShell) Run(ctx context.Context) error {
	_ = ctx
	return fmt.Errorf("platform windows: not implemented")
}
