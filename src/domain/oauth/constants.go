package oauth

const (
	queryGetAccessToken    string = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken string = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     string = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)
