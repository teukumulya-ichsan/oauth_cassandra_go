package oauth

// oauth query constant
const (
	queryGetAccessToken    string = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken string = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     string = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// oauth error const
const (
	invalidJSONBody        string = "invalid JSON body"
	invalidAccessTokenID   string = "invalid access token id"
	invalidUserID          string = "invalid user id"
	invalidClientID        string = "invalid client id"
	invalidExpireTime      string = "invalid expiration time"
	errNotFoundAccessToken string = "no access token found with given id"
)
