package usecase

import (
	"context"
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/model"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/pkg/bcrypt"
	"workshop-restful-api-backend/pkg/jwt"

	"github.com/google/uuid"
)

type IAuthUsecase interface {
	Register(ctx context.Context, param model.UserRegister) error
	Login(ctx context.Context, param model.UserLogin) (string, error)
}

type AuthUsecase struct {
	Jwt            jwt.JWT
	Bcrypt         bcrypt.IBcrypt
	UserRepository repository.IUserRepository
}

func NewAuthUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, userRepository repository.IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		Jwt:            jwt,
		Bcrypt:         bcrypt,
		UserRepository: userRepository,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, param model.UserRegister) error {
	hashedPassword, err := u.Bcrypt.GenerateHash(param.Password)
	if err != nil {
		return err
	}
	user := entity.User{
		UserId:   uuid.New(),
		Email:    param.Email,
		Username: param.Username,
		Password: hashedPassword,
		Role:     model.UserRoleUser,
	}

	err = u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecase) Login(ctx context.Context, param model.UserLogin) (string, error) {
	user, err := u.UserRepository.GetUserByUsername(ctx, param.Username)
	if err != nil {
		return "", err
	}

	err = u.Bcrypt.ValidatePassword(user.Password, param.Password)
	if err != nil {
		return "", err
	}

	token, err := u.Jwt.GenerateToken(user.UserId.String(), user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
