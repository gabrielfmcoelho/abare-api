package schemas

import (
    "gorm.io/gorm"
)


type Traits struct {
    gorm.Model
    IsValid bool
    Value   string
    ChildID uint `gorm:"foreignKey:ChildID"`
    TagID  []Tag `gorm:"foreignKey:TagID"`
}
