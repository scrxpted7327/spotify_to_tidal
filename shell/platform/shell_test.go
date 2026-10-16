package platform_test

import (
	"context"
	"errors"
	"runtime"
	"strings"
	"testing"

	"spotify_to_tidal/shell/platform"
)

func TestNewIdentifiesHostOS(t *testing.T) {
	shell := platform.New()
	if got, want := shell.OS(), runtime.GOOS; got != want {
		t.Fatalf("OS() = %q, want %q", got, want)
	}
}

func TestRunNotImplemented(t *testing.T) {
	shell := platform.New()
	err := shell.Run(context.Background())
	if err == nil {
		t.Fatal("Run() error = nil, want not-implemented error")
	}
	if !strings.Contains(err.Error(), "not implemented") {
		t.Fatalf("Run() error = %q, want substring %q", err.Error(), "not implemented")
	}
	if !strings.Contains(err.Error(), shell.OS()) {
		t.Fatalf("Run() error = %q, want OS %q mentioned", err.Error(), shell.OS())
	}
	if errors.Is(err, context.Canceled) {
		t.Fatalf("Run() unexpected context error: %v", err)
	}
}
