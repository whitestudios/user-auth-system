package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (rep *UserRepository) FindByEmail(u *User, email string) error {
	if err := rep.db.Where("email = ?", email).First(u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

func (rep *UserRepository) FindById(id uint) (*User, error) {
	var u User

	if err := rep.db.First(&u, id).Error; err != nil {
		// this error can be 'user not found, handle this when you call a user repository
		return nil, err
	}

	return &u, nil
}

func (rep *UserRepository) GetAll() ([]User, error) {
	var users []User

	if err := rep.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (rep *UserRepository) Create(u *User) error {
	return rep.db.Create(u).Error
}

func (rep *UserRepository) Update(u *User) error {
	return rep.db.Save(u).Error
}

func (rep *UserRepository) DeleteById(id uint) error {
	var user User

	if err := rep.db.First(&user, id).Error; err != nil {
		return err
	}

	return rep.db.Delete(&user).Error
}

func (rep *UserRepository) Delete(u *User) error {
	return rep.db.Delete(&u).Error
}
