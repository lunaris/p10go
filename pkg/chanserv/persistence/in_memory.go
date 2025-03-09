package persistence

import "errors"

type inMemoryUserRepository struct {
	users []InMemoryUser
}

type InMemoryUser struct {
	Username string
	Password string
}

func NewInMemoryUserRepository(users ...InMemoryUser) UserRepository {
	return &inMemoryUserRepository{
		users: users,
	}
}

func (r *inMemoryUserRepository) Authenticate(req AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	for _, u := range r.users {
		if u.Username == req.Username && u.Password == req.Password {
			return &AuthenticateUserResponse{}, nil
		}
	}

	return nil, errors.New("invalid username or password")
}
