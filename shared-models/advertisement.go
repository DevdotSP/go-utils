package sharedModels

import "time"

type Advertisement struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:text;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	URLImage    string    `gorm:"type:text;not null" json:"url_image"` // single image path
	CreatedAt   time.Time `gorm:"autoCreateTime;type:timestamptz" json:"created_at"`
	CreatedBy   string    `gorm:"type:varchar(100);not null" json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;type:timestamptz" json:"updated_at"`
	UpdatedBy   string    `gorm:"type:varchar(100)" json:"updated_by"`
}

// TableName sets the table name for Advertisement
func (Advertisement) TableName() string {
	return "v1.advertisements"
}