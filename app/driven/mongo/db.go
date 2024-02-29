package mongo

import (
	"github.com/dulguundd/logError-lib/errs"
	"restAPIServer/app/dto"
)

type Repository interface {
	Query() *errs.AppError
	QueryById() (*dto.ProductOffering, *errs.AppError)
}
