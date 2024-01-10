package service

import (
	"context"
	"user-services/constants"
	"user-services/dto"
	"user-services/model"
	"user-services/repository"
	"user-services/util"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository repository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, req *repository.ReqDeleteUser) error
	GetOneUser(ctx context.Context, req *repository.ReqFirstUser) (*model.User, error)
	GetUsers(ctx context.Context, req *repository.ReqFindAllUser, skip, limit int) ([]model.User, error)
	SignUp(ctx context.Context, req *dto.ReqSignUp) (*dto.ResSignUp, error)
	SignIn(ctx context.Context, req *dto.ReqSignIn) (*dto.ResSignUp, error)
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{UserRepository: u}
}

func (u *userService) CreateUser(ctx context.Context, user *model.User) error {
	return u.UserRepository.Create(ctx, user)
}

func (u *userService) UpdateUser(ctx context.Context, user *model.User) error {
	return u.UserRepository.Update(ctx, user)
}

func (u *userService) DeleteUser(ctx context.Context, req *repository.ReqDeleteUser) error {
	return u.UserRepository.Delete(ctx, req)
}

func (u *userService) GetOneUser(ctx context.Context, req *repository.ReqFirstUser) (*model.User, error) {
	return u.UserRepository.FindOne(ctx, req, []string{})
}

func (u *userService) GetUsers(ctx context.Context, req *repository.ReqFindAllUser, skip, limit int) ([]model.User, error) {
	return u.UserRepository.Find(ctx, req, []string{}, skip, limit)
}

func (u *userService) SignUp(ctx context.Context, req *dto.ReqSignUp) (*dto.ResSignUp, error) {
	userExist, _ := u.UserRepository.FindOne(ctx, &repository.ReqFirstUser{Email: req.Email}, []string{})
	if userExist != nil {
		return nil, constants.ERROR_EXIST
	}
	var newUser = model.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Username,
	}
	err := u.UserRepository.Create(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	token, err := util.CreateToken(map[string]interface{}{
		"id":    newUser.ID,
		"name":  newUser.Name,
		"email": newUser.Email,
	}, nil)
	if err != nil {
		return nil, err
	}
	rs := dto.ResSignUp{
		User:  newUser,
		Token: *token,
	}
	return &rs, nil
}

func (u *userService) SignIn(ctx context.Context, req *dto.ReqSignIn) (*dto.ResSignUp, error) {
	userExist, err := u.UserRepository.FindOne(ctx, &repository.ReqFirstUser{Email: req.Email}, []string{})
	if err != nil {
		return nil, constants.ERROR_NOT_FOUND
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
	if err != nil {
		return nil, constants.ERROR_SIGN_IN
	}

	token, err := util.CreateToken(map[string]interface{}{
		"id":    userExist.ID,
		"name":  userExist.Name,
		"email": userExist.Email,
	}, nil)
	if err != nil {
		return nil, err
	}
	rs := dto.ResSignUp{
		User:  *userExist,
		Token: *token,
	}
	return &rs, nil
}
