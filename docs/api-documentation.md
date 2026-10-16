# API Documentation Guide

Curated reference for playlist transfer and track matching across Spotify, TIDAL, and YouTube. This project currently integrates **Spotify** (source) and **TIDAL** (destination) via Spotipy and the third-party [tidalapi](https://tidalapi.netlify.app/) package ([repository](https://github.com/tamland/python-tidal)). YouTube is included for future or alternative workflows.

---

## Spotify Web API

**Official docs:** [Web API](https://developer.spotify.com/documentation/web-api) · [Reference](https://developer.spotify.com/documentation/web-api/reference) · [Authorization](https://developer.spotify.com/documentation/web-api/concepts/authorization) · [Scopes](https://developer.spotify.com/documentation/web-api/concepts/scopes)

### Authentication

- Uses **OAuth 2.0 Authorization Code** flow (via Spotipy `SpotifyOAuth`).
- Register an app on the Spotify Developer Dashboard and configure a redirect URI matching `config.yml`.
- **Scopes used in this repo** (`src/spotify_to_tidal/auth.py`):
  - `playlist-read-private` — read owned and collaborative playlists, including track lists.
  - `user-library-read` — read saved tracks (“Liked Songs”).

### Relevant endpoints & resources

| Operation | API resource | Notes |
|-----------|--------------|-------|
| List user playlists | `GET /me/playlists` | Paginated; filter by owner in application code. |
| Get playlist metadata | `GET /playlists/{playlist_id}` | Name, description, owner. |
| Get playlist items | `GET /playlists/{playlist_id}/items` | Current documented route; paginated and may contain tracks or episodes. |
| Get saved tracks | `GET /me/tracks` | Paginated “Liked Songs”. |
| Get current user | `GET /me` | User ID for ownership filtering. |

**Track fields used for matching** (requested via `fields` on playlist tracks): `name`, `artists`, `album` (name, artists), `track_number`, `duration_ms`, `id`, `external_ids.isrc`.

### Limitations

- **Rate limits** — throttling returns HTTP 429; the sync layer retries with backoff.
- **Pagination** — playlist and library endpoints have endpoint-specific page limits; follow the response's `next` URL until it is null.
- **No write access** — current scopes are read-only; this tool does not modify Spotify data.
- **Episodes & local files** — playlist items may be podcast episodes or unavailable tracks; the sync filters to music tracks with valid album metadata.
- **ISRC availability** — not every track exposes an ISRC in `external_ids`; matching falls back to heuristics when absent.

### Project mapping

| Code location | Spotify API usage |
|---------------|-------------------|
| `auth.py` | OAuth session with `playlist-read-private`, `user-library-read`. |
| `sync.py` → `get_playlists_from_spotify` | `current_user_playlists`, `current_user`. |
| `sync.py` → `get_tracks_from_spotify_playlist` | Spotipy's legacy-named `playlist_tracks` wrapper with selective `fields`; compare it with Spotify's current Playlist Items reference when upgrading Spotipy. |
| `sync.py` → `sync_favorites` | `current_user_saved_tracks`. |
| `sync.py` → `get_playlists_from_config` | `playlist` for explicit playlist IDs. |

---

## TIDAL

### Official API vs. `tidalapi` (this repo)

| | Official TIDAL Developer API | `tidalapi` Python package |
|--|------------------------------|---------------------------|
| **Docs** | [Overview](https://developer.tidal.com/documentation/api-sdk/api-sdk-overview) · [Authorization](https://developer.tidal.com/documentation/api-sdk/api-sdk-authorization) · [API reference](https://developer.tidal.com/apiref) | [Package docs](https://tidalapi.netlify.app/) · [GitHub](https://github.com/tamland/python-tidal) |
| **Status** | Supported developer platform | Unofficial client; reverse-engineered endpoints |
| **Used here?** | No | **Yes** — all TIDAL I/O goes through `tidalapi` |

This project does **not** call the official TIDAL Developer API directly. Behavior, rate limits, and stability depend on the unofficial `tidalapi` layer and underlying TIDAL services.

### Authentication (`tidalapi`)

- **OAuth device / browser flow** via `session.login_oauth()`; tokens persisted in `.session.yml`.
- Session reload uses `load_oauth_session` with stored access and refresh tokens.

See [official TIDAL authorization](https://developer.tidal.com/documentation/api-sdk/api-sdk-authorization) and the [tidalapi package docs](https://tidalapi.netlify.app/) for comparison.

### Relevant operations (via `tidalapi`)

| Operation | Typical usage in repo | Purpose |
|-----------|----------------------|---------|
| Search | `session.search(query, models=[Album\|Track])` | Find TIDAL tracks for Spotify source tracks. |
| List user playlists | `users/{user_id}/playlists` (via patch) | Match Spotify playlists by name. |
| Playlist tracks | `{playlist_url}/tracks` (via patch) | Read existing destination playlist. |
| Create playlist | `user.create_playlist(name, description)` | Create missing TIDAL playlists. |
| Add tracks | `playlist.add(track_ids)` | Append or rebuild playlist contents. |
| Favorites | `user.favorites` add/list | Sync “Liked Songs” (`--sync-favorites`). |

Track matching in `sync.py` compares **ISRC**, **duration**, **title**, and **artist overlap** between Spotify and TIDAL `Track` objects.

### Limitations

- **Unofficial API** — endpoints may change without notice; not covered by TIDAL developer SLAs.
- **Rate limiting** — `TooManyRequests` exceptions; sync uses concurrency limits and retries (`config.yml`: `rate_limit`, `max_concurrency`).
- **Search quality** — text search is fuzzy; album-first then track-fallback strategy may miss regional or renamed releases.
- **Catalog gaps** — tracks on Spotify may be absent or differently versioned on TIDAL.
- **No ISRC search endpoint in app code** — matching relies on search results plus local heuristics, not a dedicated ISRC lookup.

### Project mapping

| Code location | TIDAL usage |
|---------------|-------------|
| `auth.py` | `tidalapi.Session`, OAuth login, `.session.yml` persistence. |
| `sync.py` → `tidal_search` | Album search, then track search; `match()` validation. |
| `sync.py` → `sync_playlist` | Create playlist, search new tracks, update track list. |
| `sync.py` → `sync_favorites` | Favorites list + `add_track`. |
| `tidalapi_patch.py` | Paginated fetch for playlists, playlist tracks, favorites; bulk add/clear. |

---

## YouTube / YouTube Music

**Important:** Google does **not** publish a separate public **YouTube Music API**. The official supported surface is **YouTube Data API v3**, which manages YouTube **videos** and **playlists**. It does **not** guarantee YouTube Music catalog semantics (albums, artists, ISRCs, or “song-only” entities). Playlist transfer to or from YouTube Music requires mapping tracks to **video IDs** and accepting catalog mismatches.

**Official docs:** [Getting started](https://developers.google.com/youtube/v3/getting-started) · [Reference](https://developers.google.com/youtube/v3/docs) · [Authentication](https://developers.google.com/youtube/v3/guides/authentication) · [Quotas](https://developers.google.com/youtube/v3/determine_quota_cost)

### Authentication

- **OAuth 2.0** for user-owned playlists and library operations.
- **API keys** suffice for public read-only search in some cases; modifying user playlists requires OAuth.
- Enable **YouTube Data API v3** in Google Cloud Console and configure OAuth consent and credentials.

### Relevant endpoints & resources

| Operation | Resource | Quota cost (typical) |
|-----------|----------|----------------------|
| List playlists | `playlists.list` ([playlists](https://developers.google.com/youtube/v3/docs/playlists)) | 1 unit |
| List playlist items | `playlistItems.list` ([playlistItems](https://developers.google.com/youtube/v3/docs/playlistItems)) | 1 unit |
| Insert playlist item | `playlistItems.insert` ([playlistItems](https://developers.google.com/youtube/v3/docs/playlistItems)) | 50 units |
| Search (videos) | [`search.list`](https://developers.google.com/youtube/v3/docs/search/list) | 100 units |

**Playlist docs:** [Playlists](https://developers.google.com/youtube/v3/docs/playlists) · [PlaylistItems](https://developers.google.com/youtube/v3/docs/playlistItems)

### Track matching considerations

- Search returns **videos**, not normalized music tracks — match on title, channel/uploader, and duration; optional Music Brainz or ISRC metadata is not first-class in the API.
- Official **music playlist** types exist on YouTube, but item payloads still reference `videoId`.
- Duplicate uploads, lyric videos, and topic channels complicate automated matching (similar challenges to this repo’s Spotify→TIDAL heuristics).

### Limitations

- **Daily quota** — default 10,000 units/day; search-heavy sync consumes quota quickly ([quota calculator](https://developers.google.com/youtube/v3/determine_quota_cost)).
- **No YouTube Music–specific API** — cannot reliably distinguish “album track” vs. “random upload” without custom scoring.
- **Pagination** — `playlistItems.list` and `search.list` are paginated (max 50 per page).
- **Not integrated in this repo** — no YouTube client or credentials; section provided for comparison and future work.

### Project mapping

| Status | Notes |
|--------|-------|
| **Not used** | This tool syncs Spotify → TIDAL only. |
| **Analogous workflow** | Export playlist video IDs from source → `search.list` for candidates → `playlistItems.insert` on destination. |
| **Matching parallel** | Same pattern as `sync.py`: metadata comparison (title, artist/uploader, duration) when stable IDs (ISRC) are unavailable. |

---

## Cross-service comparison

| Concern | Spotify | TIDAL (via `tidalapi`) | YouTube Data API v3 |
|---------|---------|------------------------|---------------------|
| Stable track ID | Spotify track ID | TIDAL track ID | Video ID |
| ISRC in API | Often in `external_ids` | On `Track.isrc` when present | Not standard |
| Playlist read | Official, paginated | Unofficial, paginated | Official, paginated |
| Playlist write | Not used here | Unofficial | Official (quota-heavy) |
| Search for matching | N/A (source) | Text search | `search.list` (100 units/call) |
| Auth in this repo | Spotipy OAuth | `tidalapi` OAuth | — |

---

## Further reading

- [Spotify Web API](https://developer.spotify.com/documentation/web-api)
- [TIDAL Developer API overview](https://developer.tidal.com/documentation/api-sdk/api-sdk-overview)
- [tidalapi documentation](https://tidalapi.netlify.app/)
- [YouTube Data API v3](https://developers.google.com/youtube/v3/getting-started)
