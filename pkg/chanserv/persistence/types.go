package persistence

type UserRepository interface {
	Authenticate(req AuthenticateUserRequest) (*AuthenticateUserResponse, error)
}

type AuthenticateUserRequest struct {
	Username string
	Password string
}

type AuthenticateUserResponse struct{}
