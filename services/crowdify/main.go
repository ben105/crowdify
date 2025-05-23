package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ben105/crowdify/packages/db"
	"github.com/ben105/crowdify/packages/env"
)

var conn *db.DbConnection

func main() {
	conn = db.Connect()

	mux := http.NewServeMux()

	// Not a real endpoint. Just for testing.
	mux.HandleFunc("/unprocessTrack", handleUnprocessedTrack)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/callback", handleCallback)

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /login request\n")

	clientId := env.GetSpotifyClientId()
	clientSecret := env.GetSpotifyClientSecret()
	redirectUri := env.GetSpotifyRedirectUri()
	if clientId == "" || clientSecret == "" || redirectUri == "" {
		http.Error(w, "Spotify client ID and secret and redirect URI are required", http.StatusBadRequest)
		return
	}

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientId)
	params.Add("scope", scope)
	params.Add("redirect_uri", redirectUri)
	params.Add("state", randStringBytes(16))

	authUri := spotifyAuthUrl + "?" + params.Encode()
	http.Redirect(w, r, authUri, http.StatusFound)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	receivedState := r.URL.Query().Get("state")
	receivedCode := r.URL.Query().Get("code")

}

func handleUnprocessedTrack(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /unprocessTrack request\n")

	trackName := r.URL.Query().Get("name")

	// Mock Track
	mockTrack := db.UnprocessedTrack{
		ID:          randStringBytes(10),
		Name:        trackName,
		Type:        "track",
		DurationMs:  250_000,
		Popularity:  75,
		Explicit:    false,
		TrackNumber: 3,
		DiscNumber:  1,
	}

	addUnprocessedTrack(mockTrack)
	io.WriteString(w, "Success!\n")
}

func addUnprocessedTrack(unprocessedTrack db.UnprocessedTrack) {
	// Add a track to the database.
	db.InsertUnprocessedTrack(conn, &unprocessedTrack)

	// Send a message to the queue to process the track.
	trackJson, err := json.Marshal(unprocessedTrack)
	if err != nil {
		log.Fatal(err)
	}

	serverURL := strings.Join([]string{env.GetMessengerUrl(), "message"}, "/")

	bodyReader := bytes.NewReader(trackJson)
	req, err := http.NewRequest(http.MethodPost, serverURL, bodyReader)
	if err != nil {
		panic(fmt.Sprintf("Error creating HTTP request: %v", err))
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	log.Printf("Sending POST request to %s", serverURL)
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Error sending request: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		log.Println("Success! Server responded with 204 No Content as expected.")
	} else {
		log.Printf("Warning: Expected status code 204, but received %d", resp.StatusCode)

		responseBodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			log.Printf("Error reading response body: %v", readErr)
		} else if len(responseBodyBytes) > 0 {
			log.Printf("Response body content: %s", string(responseBodyBytes))
		} else {
			log.Println("Response body was empty.")
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
