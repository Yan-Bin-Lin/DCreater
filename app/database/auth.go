package database

// generate a new access token
func NewAccessToken(uid, code string) error {
	return checkAffect(db.Exec("call new_token(?, ?)", uid, code))
}

// check token with auth
func CheckAccessAuth(uid, code, oid string) (bool, error) {
	return db.SQL("call check_auth(?, ?, ?)", uid, code, oid).Exist()
}

// check token with auth
func CheckAccessToken(uid, code string) (bool, error) {
	return db.SQL("call check_token(?, ?)", uid, code).Exist()
}
