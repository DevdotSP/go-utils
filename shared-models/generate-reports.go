package sharedModels

import "time"


// this model is view table create first to use on list of export
type UserExport struct {
	UserID         string     `gorm:"column:user_id"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	MobileNumber   string     `gorm:"column:mobile_number"`
	Name           string     `gorm:"column:name"`
	Address        string     `gorm:"column:address"`
	Barangay       string     `gorm:"column:barangay"`
	Municipality   string     `gorm:"column:municipality"`
	Province       string     `gorm:"column:province"`
	Region         string     `gorm:"column:region"`
	Latitude       string     `gorm:"column:latitude"`
	Longitude      string     `gorm:"column:longitude"`
}

type UserExportRequest struct {
	ID          uint      `gorm:"primaryKey"`
	RequestDate time.Time `gorm:"autoCreateTime"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	FileName    string    `json:"file_name"`
	FilePath    string    `json:"file_path"`
	EncodedBy   string    `json:"encoded_by"`
}

func (UserExportRequest) TableName() string {
	return "v1.mobile_users_report"
}

type GenerateExportRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	EncodedBy string `json:"encoded_by"`
}
