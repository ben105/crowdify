package main

import "github.com/ben105/crowdify/packages/db"

func processTrack(unprocessedTrack *db.UnprocessedTrack) {
	// Just a stub for now, but this is where the processing logic would go.
	// For example, we will want to fetch artist data from the Spotify API.
	track := &db.Track{
		ID:          unprocessedTrack.ID,
		Name:        unprocessedTrack.Name,
		Type:        unprocessedTrack.Type,
		DurationMs:  unprocessedTrack.DurationMs,
		Popularity:  unprocessedTrack.Popularity,
		Explicit:    unprocessedTrack.Explicit,
		TrackNumber: unprocessedTrack.TrackNumber,
		DiscNumber:  unprocessedTrack.DiscNumber,
	}

	// TODO: We need to track process state, e.g.: "processing", "processed", "failed", etc.

	db.InsertTrack(conn, track)
}
