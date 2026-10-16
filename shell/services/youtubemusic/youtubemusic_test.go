package youtubemusic_test

import (
	"testing"

	"spotify_to_tidal/shell/services/domain"
	"spotify_to_tidal/shell/services/youtubemusic"
)

func TestProviderMetadata(t *testing.T) {
	info := youtubemusic.Info()
	if info.ID != "youtubemusic" {
		t.Fatalf("ID = %q, want %q", info.ID, "youtubemusic")
	}
	if info.Name != "YouTube Music" {
		t.Fatalf("Name = %q, want %q", info.Name, "YouTube Music")
	}
	if info != (domain.ProviderInfo{ID: youtubemusic.ProviderID, Name: youtubemusic.ProviderName}) {
		t.Fatalf("Info() = %#v, want constants ProviderID/ProviderName", info)
	}
}

func TestServiceEmbedsMusicService(t *testing.T) {
	var _ domain.MusicService = (youtubemusic.Service)(nil)
}
