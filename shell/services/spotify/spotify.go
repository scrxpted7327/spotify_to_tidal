package spotify

import "spotify_to_tidal/shell/services/domain"

const (
	ProviderID   = "spotify"
	ProviderName = "Spotify"
)

// Info returns stable metadata for the Spotify provider.
func Info() domain.ProviderInfo {
	return domain.ProviderInfo{ID: ProviderID, Name: ProviderName}
}

// Service is the Spotify-facing composition of the shared domain contract.
// No API client or credentials live here; this package exposes the surface only.
type Service interface {
	domain.MusicService
}
