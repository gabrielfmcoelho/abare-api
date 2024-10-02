package schemas


type FamilyChildren struct {
    ChildID  uint `gorm:"foreignKey:ChildID"`  
    FamilyID uint `gorm:"foreignKey:FamilyID"` 
}