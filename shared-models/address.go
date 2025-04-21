package sharedModels

import "time"



// Address Model.
// @swagger:model
type Address struct {
	ID               int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int          `json:"user_id" gorm:"index"`
	Active           bool         `json:"active"`
	Landmark         string       `json:"landmark"`
	Street           string       `json:"street"`
	PostalCode       string       `json:"postal_code"`
	RegionCode       string       `json:"region_code"`
	Region           Region       `gorm:"foreignKey:RegionCode;references:Code"`
	ProvinceCode     string       `json:"province_code"`
	Province         Province     `gorm:"foreignKey:ProvinceCode;references:Code"`
	MunicipalityCode string       `json:"municipality_code"`
	Municipality     Municipality `gorm:"foreignKey:MunicipalityCode;references:Code"`
	BarangayCode     string       `json:"barangay_code"`
	Barangay         Barangay     `gorm:"foreignKey:BarangayCode;references:Code"`
	CreatedAt        time.Time    `json:"created_at"`
}

func (Address) TableName() string {
	return "v1.address"
}

// Region Model.
// @swagger:model
type Region struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"type:varchar(255);unique" json:"code"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
}

func (Region) TableName() string {
	return "v1.region"
}

// Province Model.
// @swagger:model
type Province struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code       string    `gorm:"type:varchar(255);unique" json:"code"`
	RegionCode string    `gorm:"type:varchar(100)" json:"region_code"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
}

func (Province) TableName() string {
	return "v1.province"
}

// Municipality Model.
// @swagger:model
type Municipality struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code         string    `gorm:"type:varchar(255);unique" json:"code"`
	ProvinceCode string    `gorm:"type:varchar(100)" json:"province_code"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
}

func (Municipality) TableName() string {
	return "v1.municipality"
}

// Barangay Model.
// @swagger:model
type Barangay struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code             string    `gorm:"type:varchar(255);unique" json:"code"`
	MunicipalityCode string    `gorm:"type:varchar(100)" json:"municipality_code"`
	Name             string    `gorm:"type:varchar(255)" json:"name"`

}

func (Barangay) TableName() string {
	return "v1.barangay"
}