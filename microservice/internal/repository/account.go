package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MainAccGroup
type AccountGroup struct {
	gorm.Model
	ID               int
	AccountGroupCode string `gorm:"unique;not null"`
}

// AcountCode รหัสผู้ใช้งาน
// AccountGroupCode กลุ่มรหัสผู้ใช้งาน
// AccountGroupID กลุ่มของลูกค้า รวมถึงการแอดพนักงาน หรือ เจ้าของกิจการคนใหม่
type Account struct {
	gorm.Model
	ID             int       `gorm:"primaryKey;autoIncrement:false"`
	PublicID       uuid.UUID `gorm:"type:uuid"`
	AccountCode    string    `gorm:"type:varchar(20)"`
	AccountGroupID int
	AccountGroup   AccountGroup
	NamePrefixID   sql.NullString
	NamePrefix     NamePrefix
	MobilePrefix   string `gorm:"type:varchar(10)"`
	MobilePhone    string `gorm:"type:varchar(15)"`
	Email          string `gorm:"unique;not null"`
	FirstName      string `gorm:"type:varchar(25)"`
	LastName       string `gorm:"type:varchar(25)"`
	DateOfBirth    string `gorm:"type:varchar(25)"`
	GenderID       sql.NullString
	Gender         Gender
	Address        []Address `gorm:"foreignKey:RefAccountID;"`
	CountryID      int
	UserNote       string
	Partner        []Partner `gorm:"many2many:account_partner;"`
}
type Gender struct {
	ID   int
	Name string
}
type NamePrefix struct {
	ID   int
	Name string
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetOne(Account, int) (*Account, error)
	GetAll() ([]Account, error)
	Update(Account) (*Account, error)
	Delete(Account)
}
