//go:build linux

package platform

import (
	"context"
	"fmt"
)

type linuxShell struct{}

// New returns the Linux desktop shell stub.
func New() Shell {
	return linuxShell{}
}

func (linuxShell) OS() string {
	return "linux"
}

func (linuxShell) Run(ctx context.Context) error {
	_ = ctx
	return fmt.Errorf("platform linux: not implemented")
}
