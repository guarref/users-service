package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint32) (User, error)
	UpdateUserByID(id uint32, newUser User) (User, error)
	DeleteUserByID(id uint32) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserByID(id uint32) (User, error) {
	var curUser User
	result := r.db.First(&curUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	return curUser, nil
}

func (r *userRepository) UpdateUserByID(id uint32, newUser User) (User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	newresult := r.db.Model(&user).Select("email").Updates(newUser)
	if newresult.Error != nil {
		return User{}, newresult.Error
	}

	r.db.Model(&user).Find(&user, id)
	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint32) error {
	var deluser User
	result := r.db.First(&deluser, id)

	if result.Error != nil {
		return result.Error
	}

	r.db.Delete(&deluser)

	return nil
}
