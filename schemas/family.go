package schemas

import (
    "gorm.io/gorm"
)


type Family struct {
    gorm.Model
    EmergenceContactID uint `gorm:"foreignKey:UserID"` 
}