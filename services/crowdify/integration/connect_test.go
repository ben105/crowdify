package integration_tests

import (
	"testing"

	"github.com/ben105/crowdify/packages/db"
)

func TestConnection(t *testing.T) {
	// Act
	session := db.Connect()
	err := session.Query("SELECT * FROM crowdify.unprocessed_tracks").Exec()

	// Assert
	if err != nil {
		t.Errorf("Wanted to be able to run query without failure but got %v", err)
	}

}
