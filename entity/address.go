package entity

import (
	"encoding/json"
	"github.com/go-tron/example/config"
	"github.com/go-tron/local-time"
	"github.com/go-tron/snowflake-id"
	"gorm.io/gorm"
)

var addressIdWorker = snowflakeId.NewWithConfig15(config.Config)

type Address struct {
	AddressId  int64          `gorm:"primary_key" json:"addressId,string"`
	UserId     int64          `json:"userId"`
	Name       string         `json:"name"`
	Tel        string         `json:"tel"`
	Country    string         `json:"country"`
	Province   string         `json:"province"`
	City       string         `json:"city"`
	District   string         `json:"district"`
	Title      string         `json:"title"`
	Address    string         `json:"address"`
	Detail     string         `json:"detail"`
	AreaCode   string         `json:"areaCode"`
	PostalCode string         `json:"postalCode"`
	Tag        string         `json:"tag"`
	Lng        json.Number    `json:"lng"`
	Lat        json.Number    `json:"lat"`
	IsDefault  int            `json:"isDefault"`
	CreatedAt  localTime.Time `json:"createdAt"`
	UpdatedAt  localTime.Time `json:"updatedAt"`
	Deleted    int            `json:"deleted"`
	Distance   float64        `gorm:"-" json:"distance"`
	Disabled   bool           `gorm:"-" json:"disabled"`
}

func (*Address) BeforeCreate(tx *gorm.DB) error {
	id, err := addressIdWorker.NextId()
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("AddressId", id)
	return nil
}
