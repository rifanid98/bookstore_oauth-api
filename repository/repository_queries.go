package repository

const (
	GetAccessTokenQuery       = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?"
	CreateAccessTokenQuery    = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES(?,?,?,?)"
	UpdateExpirationTimeQuery = "UPDATE access_token SET expires=? WHERE access_token=?"
	UpdateQuery               = "UPDATE access_token SET user_id=?, client_id=?, expires=? WHERE access_token=? "
)
