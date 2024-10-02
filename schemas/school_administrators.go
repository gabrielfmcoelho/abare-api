package schemas

import (
    "gorm.io/gorm"
)

// SchoolAdministrators struct representa a tabela 'school-administrators'
type SchoolAdministrators struct {
    gorm.Model
    AdminUserID uint `gorm:"foreignKey:UserID"` // Chave estrangeira para User
    SchoolID    uint `gorm:"foreignKey:SchoolID"` // Chave estrangeira para School
}
