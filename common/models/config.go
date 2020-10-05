package models

import "time"

// Config maintains application's critical configurations
type Config struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Type        string    `gorm:"unique_index:type_key_UNIQUE;column:type;type:varchar(32);not null" json:"type"`
	Key         string    `gorm:"unique_index:type_key_UNIQUE;column:key;type:varchar(32);not null" json:"key"`
	Value       string    `gorm:"column:value;type:varchar(128);not null" json:"value"`
	Description string    `gorm:"column:description;type:varchar(128);not null" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}
