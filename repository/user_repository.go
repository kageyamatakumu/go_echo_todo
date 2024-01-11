package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateUserName(user * model.User, userId uint) error
	AssignUserToOrganization(user *model.User, userId uint) error
	// ログインしているユーザーの情報を取得
	GetLoggedInUserDetails(user *model.User, userId uint) error
	// 組織内のユーザー一覧情報を取得する
	GetOrganizationUsers(users *[]model.User, organizationId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepostory(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// SELECT * FROM users WHERE email = {email} ORDER BY id LIMIT 1;
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UpdateUserName(user *model.User, userId uint) error {
	// Clauses(clause.Returning{}): RETURNING句: Postgresqlの独自拡張で、insert,update,deleteの結果を返す機能。
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Update("name", user.Name)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ur *userRepository) AssignUserToOrganization(user *model.User, userId uint) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Update("organization_id", user.OrganizationId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ur *userRepository) GetLoggedInUserDetails(user *model.User, userId uint) error {
	if err := ur.db.Where("id=?", userId).Find(user).Error; err != nil {
		return err
	}

	return nil
}


func (ur *userRepository) GetOrganizationUsers(users *[]model.User, organizationId uint) error {
	if err := ur.db.Where("organization_id=?", organizationId).Find(users).Error; err != nil {
		return err
	}

	return nil
}