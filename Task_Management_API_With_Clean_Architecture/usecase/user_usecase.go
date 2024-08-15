package usecase

import (
	"net/http"
	"task_management_api_with_clean_architecture/domain"

	"errors"
	"time"
)

type IUserRepository interface {
	GetUser(username string) (domain.User, error)
	IsDatabaseEmpty() (bool, error)
	AddUser(user domain.User) error
	UpdateRole(username string) (int64, error)
}

type IPasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword string, password string) error
}

type ITokenService interface {
	GenerateToken(id string, username string, role string, expiryDuration int64) (string, error)
}

type UserUsecase struct {
	UserRepository  IUserRepository
	PasswordService IPasswordService
	TokenService    ITokenService
}

func NewUserUseCase(ur IUserRepository, ps IPasswordService, ts ITokenService) *UserUsecase {
	return &UserUsecase{UserRepository: ur, PasswordService: ps, TokenService: ts}
}

func (uu *UserUsecase) RegisterUser(user domain.User) (int, error) {
	//Check if user with this username already exists
	_, noUserErr := uu.UserRepository.GetUser(user.Username)

	if noUserErr == nil {
		return http.StatusConflict, errors.New("username already exists")
	}

	hashedPassword, hash_err := uu.PasswordService.HashPassword(user.Password)
	if hash_err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	user.Role = "user"
	user.Password = hashedPassword

	// check if the user is the first
	empty, err := uu.UserRepository.IsDatabaseEmpty()
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if empty {
		user.Role = "admin"
	}

	insert_err := uu.UserRepository.AddUser(user)

	if insert_err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}

func (uu *UserUsecase) LogUser(username string, password string) (string, int, error) {
	var user domain.User

	// Check whether the user exists or not
	user, err := uu.UserRepository.GetUser(username)

	if err != nil {
		return "", http.StatusNotFound, errors.New("invalid username or password")
	}

	//check password match
	verifyErr := uu.PasswordService.VerifyPassword(user.Password, password)
	if verifyErr != nil {
		return "", http.StatusNotFound, errors.New("invalid username or password")
	}

	//grant jwt token
	expiryDuration := time.Now().Add(time.Hour * 72).Unix()
	token, err := uu.TokenService.GenerateToken(user.ID, user.Username, user.Role, expiryDuration)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("internal server error")
	}

	return token, http.StatusOK, nil

}

func (uu *UserUsecase) Promote(username string) (int, error) {
	//check whether the user exists or not
	_, err := uu.UserRepository.GetUser(username)

	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	updated_count, err := uu.UserRepository.UpdateRole(username)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if updated_count == 0 {
		return http.StatusConflict, errors.New("user is already an admin")
	}

	return 200, nil

}
