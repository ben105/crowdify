package main

import (
	"log"

	"github.com/ben105/crowdify/packages/db"
	"github.com/google/uuid"
)

func addTrack(name string) {
	session := db.Connect()

	fakeTrackId := uuid.New()
	track := &db.UnprocessedTrack{
		ID:   fakeTrackId.String(),
		Name: name,
	}

	if err := db.InsertUnprocessedTrack(session, track); err != nil {
		log.Fatal(err)
	}
}
