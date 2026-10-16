package spotify_test

import (
	"testing"

	"spotify_to_tidal/shell/services/domain"
	"spotify_to_tidal/shell/services/spotify"
)

func TestProviderMetadata(t *testing.T) {
	info := spotify.Info()
	if info.ID != "spotify" {
		t.Fatalf("ID = %q, want %q", info.ID, "spotify")
	}
	if info.Name != "Spotify" {
		t.Fatalf("Name = %q, want %q", info.Name, "Spotify")
	}
	if info != (domain.ProviderInfo{ID: spotify.ProviderID, Name: spotify.ProviderName}) {
		t.Fatalf("Info() = %#v, want constants ProviderID/ProviderName", info)
	}
}

func TestServiceEmbedsMusicService(t *testing.T) {
	var _ domain.MusicService = (spotify.Service)(nil)
}
