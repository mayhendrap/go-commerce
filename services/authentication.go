package services

import (
	"go-commerce/dtos/request"
	"go-commerce/entities"
	"go-commerce/repositories"
	"go-commerce/token"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticationService interface {
	Register(req request.RegisterRequest) (entities.User, error)
	Login(req request.LoginRequest) (string, error)
}

type authenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) *authenticationService {
	return &authenticationService{userRepository: userRepository}
}

func (a *authenticationService) Register(req request.RegisterRequest) (entities.User, error) {
	var err error
	var user entities.User

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return user, err
	}

	user = entities.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashedPassword),
	}
	user, err = a.userRepository.Create(user)
	if err != nil {
		return user, err
	}
	return user, err
}

func (a *authenticationService) Login(req request.LoginRequest) (string, error) {
	user, err := a.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	maker := token.NewPasetoMaker()
	signedToken, errToken := maker.CreateToken(user.ID.String(), time.Minute*5)
	if errToken != nil {
		return "", errToken
	}

	return signedToken, nil
}
