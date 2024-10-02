package schemas

import (
    "time"
    "gorm.io/gorm"
)


type Child struct {
    gorm.Model
    Name          string
    BirthDate     time.Time `gorm:"type:date"` 
    Gender        string
    SpectrumDegree string
    TraitID       string
}
