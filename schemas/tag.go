package schemas

import (
    "gorm.io/gorm"
)


type Tag struct {
    gorm.Model
    Value string
}
