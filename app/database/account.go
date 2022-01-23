package database

func Login(userName string, password string) ([]UserOut, error) {
	userDatas := make([]UserOut, 0)
	err := db.SQL("call login(?, ?)", userName, password).Find(&userDatas)
	if err != nil {
		return nil, err
	} else if len(userDatas) == 0 {
		return nil, nil
	}

	return userDatas, nil
}

// return user page
func GetUser(uid int) (*UserOut, error) {
	userData := &UserOut{}
	has, err := db.SQL("call get_user(?)", uid).Get(userData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return userData, nil
}

// insert an user if uid is 0 else update
func PutUser(uid, username, password, email, salt string) error {
	return checkAffect(db.Exec("call put_user(?, ?, ?, ?, ?)", uid, username, password, email, salt))
}

// delete an user
func DelUser(uid, username, password string) error {
	return checkAffect(db.Exec("call del_user(?, ?, ?)", uid, username, password))
}

func GetSalt(userName string) (string, error) {
	var salt string
	has, err := db.SQL("call get_user_salt(?)", userName).Get(&salt)
	if err != nil {
		return "", err
	} else if !has {
		return "", nil
	}
	return salt, nil
}
