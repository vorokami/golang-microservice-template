package userrepo

import (
	"context"
	"golang-microservice-template/database"
	"golang-microservice-template/logger"
)

type UserRepo struct {
	dbFactory     database.DbFactory
	loggerFactory logger.LoggerFactory
}

func New(dbFactory database.DbFactory, loggerFactory logger.LoggerFactory) *UserRepo {
	return &UserRepo{dbFactory, loggerFactory}
}

var ctx = context.Background()

const recordNotFoundError = "record not found"
