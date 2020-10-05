package user

import (
	"github.com/jinzhu/gorm"

	"tazapay.com/elearning/common/models"
)

// Gorm database interaction with gorm
type Gorm struct {
	db *gorm.DB
}

// getObj return operable orm object
func (g *Gorm) getObj() *gorm.DB {
	return g.db
}

// SetObj set operable orm object
func (g *Gorm) SetObj(i interface{}) {
	g.db = i.(*gorm.DB)
}

// Create add new user
func (g *Gorm) Create(user *models.User) error {
	return g.getObj().Create(user).Error
}
