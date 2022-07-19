package usecase

import (
	"context"
	"errors"
	"simpson/config"
	"simpson/internal/service"
)

type Usecase struct {
	UserUsecase       UserUsecase
	PartnerUsecase    PartnerUsecase
	JwtUsecase        JwtUsecase
	PermissionUsecase PermissionUsecase
}

func InitUsecase(ctx context.Context, repo service.Service, cfg *config.Config) (*Usecase, error) {
	if repo == nil {
		return nil, errors.New("repo empty in usecase")
	}
	pri, pub, sign, err := ParseKey(cfg)
	if err != nil {
		return nil, err
	}
	jwtUsecase := NewJwtUsecase(cfg, pri, pub, sign)
	userUsecase := NewUserUsecase(cfg, repo.NewUserService(), jwtUsecase)
	permissionUsecase := NewPermissionUsecase(cfg, repo.NewPermissionService())
	return &Usecase{
		UserUsecase:       userUsecase,
		PartnerUsecase:    NewPartnerUsecase(repo.NewPartnerService()),
		JwtUsecase:        jwtUsecase,
		PermissionUsecase: permissionUsecase,
	}, nil
}
