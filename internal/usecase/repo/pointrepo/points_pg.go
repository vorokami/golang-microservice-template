package pointrepo

import "golang-microservice-template/database"

type PointRepo struct {
	db database.DbFactory
}

func New(db database.DbFactory) *PointRepo {
	return &PointRepo{db}
}
