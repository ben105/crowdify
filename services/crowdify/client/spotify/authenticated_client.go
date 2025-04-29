package spotify

type AuthenticatedSpotifyClient struct {
	accessToken       string
	refreshToken      string
	accessTokenExpiry int64
}

type SpotifyUser struct {
	userId      string
	displayName string
	email       string
}

func (s *AuthenticatedSpotifyClient) GetUser() string {

}
