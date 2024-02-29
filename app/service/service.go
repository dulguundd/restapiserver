package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/dulguundd/logError-lib/logger"
	"restAPIServer/app/driven/mongo"
)

type Service interface {
	MongoList() *errs.AppError
	MongoById() *errs.AppError
}

type DefaultService struct {
	repo mongo.Repository
}

func (s DefaultService) MongoList() *errs.AppError {
	_ = s.repo.Query()
	return nil
}

func (s DefaultService) MongoById() *errs.AppError {
	logger.Info("Id Query worked")
	_ = s.repo.QueryById()
	return nil
}

func NewService(repository mongo.Repository) DefaultService {
	return DefaultService{repository}
}
