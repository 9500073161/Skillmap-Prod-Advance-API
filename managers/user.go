package managers

import (
	"errors"

	"github.com/9500073161/skill-map-prod/common"
	"github.com/9500073161/skill-map-prod/models"
	"github.com/9500073161/skill-map-prod/storage"
)

type UserManager interface {
	Create(userData *common.UserCreationInput) (*models.User, error)
	List() ([]models.User, error)
	Get(id string) (models.User, error)
	Update(userId string, userData *common.UserUpdateInput) (*models.User, error)
	Delete(id string) error
}

type userManager struct {
	// DatabaseDriver
	// dbClient
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userMgr *userManager) Create(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	storage.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newUser, nil
}

func (userMgr *userManager) List() ([]models.User, error) {
	users := []models.User{}
	storage.DB.Find(&users)
	return users, nil
}

func (userMgr *userManager) Get(id string) (models.User, error) {
	user := models.User{}

	storage.DB.First(&user, id)

	return user, nil
}

func (userMgr *userManager) Update(userId string, userData *common.UserUpdateInput) (*models.User, error) {

	user := models.User{}

	storage.DB.First(&user, userId)
	storage.DB.Model(&user).Updates(models.User{FullName: userData.FullName, Email: userData.Email})

	return &user, nil
}

func (userMgr *userManager) Delete(id string) error {
	user := models.User{}

	storage.DB.First(&user, id)
	storage.DB.Delete(&user)
	return nil
}
