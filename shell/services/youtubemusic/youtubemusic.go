package youtubemusic

import "spotify_to_tidal/shell/services/domain"

const (
	ProviderID   = "youtubemusic"
	ProviderName = "YouTube Music"
)

// Info returns stable metadata for the YouTube Music provider.
func Info() domain.ProviderInfo {
	return domain.ProviderInfo{ID: ProviderID, Name: ProviderName}
}

// Service is the YouTube Music-facing composition of the shared domain contract.
// No API client or credentials live here; this package exposes the surface only.
type Service interface {
	domain.MusicService
}
