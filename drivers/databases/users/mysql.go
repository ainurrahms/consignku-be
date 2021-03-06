package users

import (
	"consignku/bussiness/users"
	"context"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (nr *mysqlUsersRepository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) GetByUsernamePassword(ctx context.Context, username, password string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("username = ? AND password = ?", username, password).First(&rec).Error

	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) Store(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *mysqlUsersRepository) FindByID(id int) (users.Domain, error) {
	rec := Users{}

	if err := nr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}
