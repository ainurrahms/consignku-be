package users

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"consignku/bussiness/indonesia_city_location"
	"consignku/helper/encrypt"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository        Repository
	contextTimeout        time.Duration
	jwtAuth               *middleware.ConfigJWT
	IndonesiaCityLocation indonesia_city_location.Repository
}

func NewUserUseCase(ur Repository, jwtauth *middleware.ConfigJWT, ic indonesia_city_location.Repository, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository:        ur,
		jwtAuth:               jwtauth,
		IndonesiaCityLocation: ic,
		contextTimeout:        timeout,
	}
}

func (uc *userUsecase) Register(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)

	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}

	id, err := uc.IndonesiaCityLocation.GetCityLocation(ctx, userDomain.IDCity)

	if err != nil {
		log.Default().Printf("%+v", err)
	}
	jsonMarshal, err := json.Marshal(id)
	if err != nil {
		log.Default().Printf("%+v", err)
	}
	userDomain.CityName = string(jsonMarshal)

	if existedUser != (Domain{}) {
		return bussiness.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)

	if err != nil {
		return bussiness.ErrInternalServer
	}

	err = uc.userRepository.Store(ctx, userDomain)

	if err != nil {
		return err
	}

	return nil
}

func (uc *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", bussiness.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", bussiness.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (uc *userUsecase) GetByID(ctx context.Context, id int) (Domain, error) {

	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.userRepository.FindByID(id)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
