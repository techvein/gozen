package oauth

type User interface {
	GenerateLoginUrl() string
	GetID() *int
	GetName() *string
	GetEmail() *string
	GetSource() string
	GetClientID() *string
	GetClientSecret() *string
	Callback(state string, code string) (User, error)
}
