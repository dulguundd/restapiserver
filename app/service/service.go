package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/dulguundd/logError-lib/logger"
	"restAPIServer/app/driven/mongo"
	"restAPIServer/app/dto"
)

type Service interface {
	MongoList() *errs.AppError
	MongoById() (dto.ProductOffering, *errs.AppError)
}

type DefaultService struct {
	repo mongo.Repository
}

func (s DefaultService) MongoList() *errs.AppError {
	_ = s.repo.Query()
	return nil
}

func (s DefaultService) MongoById() (dto.ProductOffering, *errs.AppError) {
	logger.Info("This is service level log")
	result, _ := s.repo.QueryById()
	return result, nil
}

func NewService(repository mongo.Repository) DefaultService {
	return DefaultService{repository}
}
