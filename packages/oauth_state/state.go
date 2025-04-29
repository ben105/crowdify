package oauth_state

type OAuthState interface {
	GetState(key string) string
	SetState(key string, value string)
}

type OAuthStateManager struct {
	state map[string]string
}

func NewOAuthStateManager() *OAuthStateManager {
	return &OAuthStateManager{
		state: make(map[string]string),
	}
}

func (o *OAuthStateManager) GetState(key string) string {
	return o.state[key]
}

func (o *OAuthStateManager) SetState(key string, value string) {
	o.state[key] = value
}
