package mongo

import "github.com/dulguundd/logError-lib/errs"

type Repository interface {
	Query() *errs.AppError
	QueryById() *errs.AppError
}
