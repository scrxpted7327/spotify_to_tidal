package tidal

import "spotify_to_tidal/shell/services/domain"

const (
	ProviderID   = "tidal"
	ProviderName = "TIDAL"
)

// Info returns stable metadata for the TIDAL provider.
func Info() domain.ProviderInfo {
	return domain.ProviderInfo{ID: ProviderID, Name: ProviderName}
}

// Service is the TIDAL-facing composition of the shared domain contract.
// No API client or credentials live here; this package exposes the surface only.
type Service interface {
	domain.MusicService
}
