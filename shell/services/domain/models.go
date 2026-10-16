package domain

// ProviderInfo identifies a music service provider.
type ProviderInfo struct {
	ID   string
	Name string
}

// Artist is a canonical artist identity across providers.
type Artist struct {
	ID   string
	Name string
}

// Album is a canonical album identity across providers.
type Album struct {
	ID   string
	Name string
}

// Track is a canonical track across providers.
type Track struct {
	ID       string
	Title    string
	Artists  []Artist
	Album    Album
	Duration int // milliseconds; 0 if unknown
	ISRC     string
}

// Playlist is a canonical playlist across providers.
type Playlist struct {
	ID          string
	Name        string
	Description string
	Tracks      []Track
}
