package service

import (
	"github.com/Matheus-Lara/orare/internal/api/dto"
	"github.com/Matheus-Lara/orare/internal/api/errors"
	"github.com/Matheus-Lara/orare/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (service *UserService) Profile(c *gin.Context) (*dto.UserDTO, *errors.Error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, errors.NewError("User ID not found in context", nil)
	}

	id, ok := userID.(uint)
	if !ok {
		return nil, errors.NewError("Invalid user ID type in context", nil)
	}

	user, err := service.UserRepo.FindByID(id)
	if err != nil || user == nil {
		return nil, errors.NewError("User not found", nil)
	}

	return user.ToUserDTO(), nil
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}
