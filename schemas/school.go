package schemas

import (
	"gorm.io/gorm"
)

type Schools struct {
	gorm.Model
	Name             string
	Address          string
	ContactEmail     string
	ContactTelephone string
	Children         []Child
}
