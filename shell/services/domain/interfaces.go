package domain

import "context"

// Authenticator establishes and reports provider session state.
// Implementations must not live in this package; this is the contract only.
type Authenticator interface {
	Authenticate(ctx context.Context) error
	Authenticated(ctx context.Context) (bool, error)
}

// PlaylistReader lists and loads playlists from a provider (source).
type PlaylistReader interface {
	ListPlaylists(ctx context.Context) ([]Playlist, error)
	GetPlaylist(ctx context.Context, id string) (Playlist, error)
}

// PlaylistWriter creates and mutates playlists on a provider (sink).
type PlaylistWriter interface {
	CreatePlaylist(ctx context.Context, name, description string) (Playlist, error)
	AddTracks(ctx context.Context, playlistID string, tracks []Track) error
}

// TrackMatcher resolves a canonical track to a provider-specific match.
type TrackMatcher interface {
	MatchTrack(ctx context.Context, track Track) (Track, error)
}

// MusicService composes the full provider contract used by sync tooling.
type MusicService interface {
	Authenticator
	PlaylistReader
	PlaylistWriter
	TrackMatcher
	Provider() ProviderInfo
}
