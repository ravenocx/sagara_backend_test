package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Clothes struct {
	ID    string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Color string  `json:"color" gorm:"type:varchar(50);notnull"`
	Size  string  `json:"size" gorm:"type:clothes_size;notnull"`
	Price float64 `json:"price" gorm:"type:numeric(10,2);notnull"`
	Stock int     `json:"stock" gorm:"type:int;notnull"`
}

func (c *Clothes) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}
