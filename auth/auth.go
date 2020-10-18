package auth

type UserInfo struct {
	Username string
	Name string
	Email string
	StuNumber string
	Department string
}

func Authenticate(username, password string) (*UserInfo, error){
	rawUserInfo, err := headless(username, password)
	if err != nil{return nil, err}
	user := parse(rawUserInfo)
	return user, nil
}