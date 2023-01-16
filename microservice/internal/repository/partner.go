package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	ID              int
	PartnerCode     string
	PartnerPublicID uuid.UUID `gorm:"type:uuid"`
	AccountGroupID  int
	AccountGroup    AccountGroup
	PartnerBranch   []PartnerBranch `gorm:"foreignKey:RefPartnerID;"`
	Address         []Address       `gorm:"foreignKey:RefPartnerID;"`
}

type PartnerBranch struct {
	ID           int
	RefPartnerID int
	BranchName   string
	BranchNumber string    `gorm:"type:varchar(100)"`
	Address      []Address `gorm:"foreignKey:RefPartnerBranchID;"`
}
