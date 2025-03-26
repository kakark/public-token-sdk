package tokentype

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
	AppToken     TokenType = "app_token"
	IDToken      TokenType = "id_token"
)
