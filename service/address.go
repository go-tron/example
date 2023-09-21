package service

import (
	"github.com/go-tron/example/config"
	"github.com/go-tron/example/entity"
	"github.com/go-tron/mysql"
)

var AddressService = addressService{
	&mysql.BaseService{
		DB:    config.DB,
		Model: &entity.Address{},
	},
}

type addressService struct {
	*mysql.BaseService
}
