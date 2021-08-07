package app

import (
	"log"

	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/graph/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *model.User) (*model.User, error)
}

type UserRespositoryDb struct {
	Db *gorm.DB
}

func (repo UserRespositoryDb) Insert(user *model.User) (*model.User, error) {

	err := repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error perciste %v", err)
		return user, err
	}

	return user, nil

}
