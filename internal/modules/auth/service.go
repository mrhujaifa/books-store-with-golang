package auth

type Service interface {
	SyncUserFromAuth0(claims Auth0UserClaims) (*User, error)
	GetUserByID(id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Auth0 থেকে claims পেয়ে DB-তে user create/update করবে.
// এটা better-auth এর "user sync" feel এর মতো.
func (s *service) SyncUserFromAuth0(claims Auth0UserClaims) (*User, error) {
	user, err := s.repo.FindByAuth0ID(claims.Sub)
	if err != nil {
		if IsNotFound(err) {
			newUser := &User{
				Auth0ID: claims.Sub,
				Email:   claims.Email,
				Name:    claims.Name,
				Picture: claims.Picture,
			}
			if err := s.repo.Create(newUser); err != nil {
				return nil, err
			}
			return newUser, nil
		}
		return nil, err
	}

	user.Email = claims.Email
	user.Name = claims.Name
	user.Picture = claims.Picture

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}
