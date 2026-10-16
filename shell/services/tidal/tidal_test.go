package tidal_test

import (
	"testing"

	"spotify_to_tidal/shell/services/domain"
	"spotify_to_tidal/shell/services/tidal"
)

func TestProviderMetadata(t *testing.T) {
	info := tidal.Info()
	if info.ID != "tidal" {
		t.Fatalf("ID = %q, want %q", info.ID, "tidal")
	}
	if info.Name != "TIDAL" {
		t.Fatalf("Name = %q, want %q", info.Name, "TIDAL")
	}
	if info != (domain.ProviderInfo{ID: tidal.ProviderID, Name: tidal.ProviderName}) {
		t.Fatalf("Info() = %#v, want constants ProviderID/ProviderName", info)
	}
}

func TestServiceEmbedsMusicService(t *testing.T) {
	var _ domain.MusicService = (tidal.Service)(nil)
}
