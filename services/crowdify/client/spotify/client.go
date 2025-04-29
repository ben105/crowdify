package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ben105/crowdify/packages/env"
)

type SpotifyClient struct {
	clientId     string
	clientSecret string
	redirectUri  string
}

func NewSpotifyClient(clientId, clientSecret, redirectUri string) *SpotifyClient {
	return &SpotifyClient{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectUri:  redirectUri,
	}
}

type HttpError struct {
	Message string
	Code    int
}

func (s *SpotifyClient) GetAuthCode(state string) {

}

func (s *SpotifyClient) GetAccessToken(authCode, state string) HttpError {
	if state == "" {
		return errors.New("Authorization state missing"), http.StatusBadRequest
	}

	if authCode == "" {
		return errors.New("Authorization code missing"), http.StatusBadRequest
	}

	formData := url.Values{}
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", authCode)
	formData.Set("redirect_uri", env.GetSpotifyRedirectUri())
	req, err := http.NewRequestWithContext(context.Background(), "POST", SpotifyTokenUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Printf("Error creating token request: %v", err)
		return errors.New("Internal Server Error"), http.StatusInternalServerError
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := base64.StdEncoding.EncodeToString([]byte(env.GetSpotifyClientId() + ":" + env.GetSpotifyClientSecret()))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to Spotify token endpoint: %v", err)
		http.Error(w, "Failed to contact authentication server", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading token response body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Spotify token endpoint returned error %d: %s", resp.StatusCode, string(body))
		http.Error(w, fmt.Sprintf("Spotify error: %s", string(body)), resp.StatusCode) // Pass along Spotify's error message if possible
		return
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		log.Printf("Error parsing JSON token response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully retrieved tokens. Access Token expires in: %d", tokenResp.ExpiresIn)
}
