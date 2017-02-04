package json

type UserJson struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

func NewUserJson() *UserJson {
	return new(UserJson)
}
