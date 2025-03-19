package db

type Track struct {
	ID          string `json:"id" cql:"id"`
	Name        string `json:"name" cql:"name"`
	Type        string `json:"type" cql:"type"`
	DurationMs  int    `json:"duration_ms" cql:"duration_ms"`
	Popularity  int    `json:"popularity" cql:"popularity"`
	Explicit    bool   `json:"explicit" cql:"explicit"`
	TrackNumber int    `json:"track_number" cql:"track_number"`
	DiscNumber  int    `json:"disc_number" cql:"disc_number"`
}

func InsertTrack(conn *DbConnection, track *Track) error {
	if err := conn.session.Query(`
		INSERT INTO tracks (
			id, 
			name, 
			type, 
			duration_ms, 
			popularity, 
			explicit, 
			track_number, 
			disc_number
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		track.ID,
		track.Name,
		track.Type,
		track.DurationMs,
		track.Popularity,
		track.Explicit,
		track.TrackNumber,
		track.DiscNumber,
	).Exec(); err != nil {
		return err
	}
	return nil
}

type UnprocessedTrack struct {
	ID          string `json:"id" cql:"id"`
	Name        string `json:"name" cql:"name"`
	Type        string `json:"type" cql:"type"`
	DurationMs  int    `json:"duration_ms" cql:"duration_ms"`
	Popularity  int    `json:"popularity" cql:"popularity"`
	Explicit    bool   `json:"explicit" cql:"explicit"`
	TrackNumber int    `json:"track_number" cql:"track_number"`
	DiscNumber  int    `json:"disc_number" cql:"disc_number"`

	// Album information
	AlbumID                   string `json:"album_id" cql:"album_id"`
	AlbumName                 string `json:"album_name" cql:"album_name"`
	AlbumType                 string `json:"album_type" cql:"album_type"`
	AlbumTotalTracks          int    `json:"album_total_tracks" cql:"album_total_tracks"`
	AlbumReleaseDate          string `json:"album_release_date" cql:"album_release_date"`
	AlbumReleaseDatePrecision string `json:"album_release_date_precision" cql:"album_release_date_precision"`
}

func GetUnprocessedTrack(conn *DbConnection, id string) (*UnprocessedTrack, error) {
	track := &UnprocessedTrack{}
	if err := conn.session.Query(`SELECT * FROM unprocessed_tracks WHERE id = ?`, id).Scan(
		&track.ID,
		&track.Name,
		&track.Type,
		&track.DurationMs,
		&track.Popularity,
		&track.Explicit,
		&track.TrackNumber,
		&track.DiscNumber,
		&track.AlbumID,
		&track.AlbumName,
		&track.AlbumType,
		&track.AlbumTotalTracks,
		&track.AlbumReleaseDate,
		&track.AlbumReleaseDatePrecision,
	); err != nil {
		return nil, err
	}
	return track, nil
}

func InsertUnprocessedTrack(conn *DbConnection, track *UnprocessedTrack) error {
	if err := conn.session.Query(`
		INSERT INTO unprocessed_tracks (
			id, 
			name, 
			type, 
			duration_ms, 
			popularity, 
			explicit, 
			track_number, 
			disc_number, 
			album_id, 
			album_name, 
			album_type, 
			album_total_tracks, 
			album_release_date, 
			album_release_date_precision
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		track.ID,
		track.Name,
		track.Type,
		track.DurationMs,
		track.Popularity,
		track.Explicit,
		track.TrackNumber,
		track.DiscNumber,
		track.AlbumID,
		track.AlbumName,
		track.AlbumType,
		track.AlbumTotalTracks,
		track.AlbumReleaseDate,
		track.AlbumReleaseDatePrecision,
	).Exec(); err != nil {
		return err
	}
	return nil
}
