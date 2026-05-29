package auth

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindByAuth0ID(auth0ID string) (*User, error)
	FindByID(id uint) (*User, error)
	Create(user *User) error
	Update(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByAuth0ID(auth0ID string) (*User, error) {
	var user User
	err := r.db.Where("auth0_id = ?", auth0ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
