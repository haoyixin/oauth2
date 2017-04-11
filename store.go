package oauth2

type (
	// ClientStore the client information storage interface
	ClientStore interface {
		// create and store new client information
		Create(cli ClientInfo) error

		// use the ID to delete the client information
		RemoveByID(id string) error

		// user the ID for client information data
		GetByID(id string) (ClientInfo, error)
	}

	// TokenStore the token information storage interface
	TokenStore interface {
		// create and store the new token information
		Create(info TokenInfo) error

		// delete the authorization code
		RemoveByCode(code string) error

		// use the access token to delete the token information
		RemoveByAccess(access string) error

		// use the refresh token to delete the token information
		RemoveByRefresh(refresh string) error

		// use the authorization code for token information data
		GetByCode(code string) (TokenInfo, error)

		// use the access token for token information data
		GetByAccess(access string) (TokenInfo, error)

		// use the refresh token for token information data
		GetByRefresh(refresh string) (TokenInfo, error)
	}
)
